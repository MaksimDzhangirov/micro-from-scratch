# Очередь задач с помощью Go и RabbitMQ

Мы хотим построить систему очереди задач. Мы хотим поддерживать очередь 
задач/работ, куда мы будем добавлять новую задачу. Workers будут активно 
следить за очередью и выполнять эту задачу по мере ее поступления. Система 
очередей с задачами идеально подходит для фоновых заданий, которые могут занимать 
больше времени, чем ваш средний HTTP-запрос. Типичный пример, который я часто 
использую для описания очередей задач - если ваше приложение обрабатывает 
фотографии, загруженные пользователями, создает несколько версий (миниатюры, 
разные размеры) и делится ими в разных социальных сетях, как бы вы справились 
с этим? Изменение размера фотографий и загрузка их на другие сайты определенно 
займет время. Вы хотите сделать это внутри вашего обработчика http? Так делать 
не следует. Вместо этого, когда фотография загружена, сохраните фотографию в 
удобном месте и передайте детали фоновым workers, которые затем обработают 
фотографию и загрузят ее туда, где это необходимо.

## Нам нужна система обмена сообщениями

Чтобы осуществить это, нам нужна система, которая может работать как очередь 
сообщений. Мы хотим иметь распределенную систему очередей задач, в которой
исполнители и создатели задач будут находиться на разных серверах. Может быть 
множество источников задач и множество workers, обрабатывающих их. Для построения 
такой распределенной системы нам нужна централизованная очередь сообщений/брокер 
сообщений. Нам нужна система, в которой мы можем передавать сообщения, и эта 
система будет доставлять эти сообщения нашим workers.

Есть несколько инструментов, которые могут нам в этом помочь — Redis, Kafka, 
RabbitMQ, ZeroMQ, IronMQ, AWS SQS — выбор велик. Redis очень популярен, и лично 
я часто использую его с Celery/Bull. Но в некоторых аспектах RabbitMQ — лучший 
выбор. В нашем примере мы собираемся использовать RabbitMQ с Go для создания 
очень простой системы очередей задач.

## Понимание понятий RabbitMQ

Существует несколько понятий, о которых следует помнить при работе с RabbitMQ.

**Производители и потребители:** Производитель — это тот, кто создает новые 
сообщения / задачу. Потребители – это те, кто их потребляет. В нашем примере, 
когда происходит загрузка файла, обработчик http создает сообщение для обработки 
нашими workers. Обработчик http / веб-приложение является производителем, а 
фоновые workers — потребителями.

**Обменник и очереди:** по названию можно догадаться, что они каким-то образом связаны 
с обработкой сообщений. Обменники получают сообщения от производителей и доставляют 
их в очереди. Потребители потребляют сообщения из очередей. RabbitMQ 
поставляется с очень мощными функциями маршрутизации сообщений. Мы можем разными 
способами настроить способ доставки сообщений в разные очереди.

## Пример с командной строкой

В этом посте мы действительно не хотим заниматься созданием полноценного 
веб-приложения с загрузкой файлов. Мы хотим реализовать что-то по-быстрому.
Поэтому мы будем делать это в командной строке. Мы создадим инструмент публикации
в командной строке, который будет публиковать/создавать сообщения. И 
потребитель, который будет потреблять сообщения. Затем мы запустим несколько 
экземпляров потребителей параллельно, чтобы показать, как мы можем 
масштабировать эту систему, добавляя больше рабочих процессов.

Мы собираемся создать очень сложный (?!?) калькулятор, который может принимать 
два числа и выводить их сумму на стандартный вывод. И нам нужно сделать его 
веб-масштабируемым, поэтому нам нужно использовать Go и RabbitMQ.

## Начальная настройка

Прежде чем мы сможем начать создавать калькулятор на миллион долларов, нам 
нужно где-то установить RabbitMQ, к которому мы можем подключиться. Для нас, 
разработчиков, что может быть лучше, чем localhost? Давайте установим RabbitMQ 
на наш локальный компьютер и запустим его.

Установка RabbitMQ может варьироваться от платформы к платформе. На MacBook я 
установил его с помощью Homebrew. В дистрибутиве Linux он, вероятно, доступен из 
менеджера пакетов. Для Windows должны быть устанавливаемые пакеты.

После того, как мы установили RabbitMQ, нам нужно установить пакет Go AMQP в 
нашей системе. Я использую Go Modules для его установки. Вы можете выполнить
команду `go get` или использовать систему управления зависимостями.

```shell
go get github.com/streadway/amqp
```

## Создание потребителя

Мы создали каталог `consumer`, внутри которого мы будем создавать наше 
потребительское приложение. Потребительское приложение должно подключиться к 
RabbitMQ, объявить очередь, которую оно хочет прослушивать, а затем начать 
потреблять сообщения.

Прежде чем мы начнем, мы собираемся создать функцию обработки ошибок. Это должно 
помочь нам в утомительном сценарии обработки ошибок Go.

```go
func handleError(err error, msg string) {
    if err != nil {
        log.Fatal("%s: %s", msg, err)	
    }
}
```

Если ошибка не равна `nil`, выводим сообщение, подробности об ошибке и завершаем 
работу приложения. Это то, что делает вышеприведенная функция.

Теперь приступим к настройке подключения к RabbitMQ.

```go
amqpHost := os.Getenv("AMQP_HOST")
amqpPort := os.Getenv("AMQP_PORT")
amqpUser := os.Getenv("AMQP_USER")
amqpPassword := os.Getenv("AMQP_PASSWORD")
amqpConnectionURL := fmt.Sprintf("amqp://%s:%s@%s:%s", amqpUser, amqpPassword, amqpHost, amqpPort)
conn, err := amqp.Dial(amqpConnectionURL)
handleError(err, "Can't connect to AMQP")
defer conn.Close()
```

Мы пробуем подсоединиться к RabbitMQ и завершаем работу, если у нас не 
получилось. Настройки подключения хранятся в файле `.env.dev`.

```shell
AMQP_HOST=rabbitmq.client.local
AMQP_PORT=5672
AMQP_USER=guest
AMQP_PASSWORD=guest
```

Если соединение установлено, нам нужно установить канал. Не путайте его с 
каналом Go. RabbitMQ имеет собственное понятие каналов. Соединение — это 
TCP-соединение от клиента к серверу. Соединение нелегко создать. Канал служит 
протоколом связи по соединению. Каналы достаточно просто создать. Мы должны 
стремиться ограничить количество подключений до минимального числа, создав 
столько каналов, сколько нам нужно, поверх этих подключений.

```go
amqpChannel, err := conn.Channel()
handleError(err, "Can't create a amqpChannel")

defer amqpChannel.Close()
```

Теперь мы можем начать взаимодействовать с RabbitMQ. Нам нужно сообщить серверу 
об интересующей нас очереди.

```go
queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
handleError(err, "Could not declare `add` queue")

err = amqpChannel.Qos(1, 0, false)
handleError(err, "Could not configure QoS")
```

Затем мы создаём очередь с именем `add`. Если вам интересно что представляют 
собой аргументы функции взгляните на [документацию здесь](https://godoc.org/github.com/streadway/amqp#Channel.QueueDeclare).

RabbitMQ начинает доставлять сообщения потребителям в циклическом режиме.
Таким образом, он поровну распределяет работу между всеми workers. Если какая-то 
работа будет длиться дольше, а какая-то закончится первой, у одного worker 
накопится много задач, а другой вообще не вспотеет. Один worker всегда будет 
занят, другой всегда будет бездействовать. Чтобы исключить такие сценарии, мы 
просим RabbitMQ доставлять новые сообщения только тогда, когда worker
подтвердил предыдущее сообщение. В [документации](https://godoc.org/github.com/streadway/amqp#Channel.Qos) по функциям Qos можно найти более 
подробное объяснение.

Давайте продолжим и начнем потреблять сообщения.

```go
messageChannel, err := amqpChannel.Consume(
    queue.Name,
    "",
    false,
    false,
    false,
    false,
    nil,
)
handleError(err, "Could not register consumer")
```

Посмотреть что представляют собой аргументы можно [здесь](https://godoc.org/github.com/streadway/amqp#Channel.Consume).
В этот раз мы получим go канал. Мы можем использовать `range`, чтобы пройтись по 
элементам этого канала для получения сообщений.

Мы хотим отправлять сообщения в формате JSON. Чтобы представить задачу для 
операции добавления, мы должны снова определить тип в нашем файле `shared.go`.

```go
type addTask struct {
    Number1 int
    Number2 int
}
```

Имея `messageChannel` мы можем начать анализировать его, декодируя тела сообщений 
в экземпляры `AddTask`, а затем суммируя `Number1` и `Number2`, чтобы получить 
результат.

```go
stopChan := make(chan bool)

go func() {
    log.Printf("Consumer ready, PID: %d", os.Getpid())
    for d := range messageChannel {
        log.Printf("Received a message: %s", d.Body)

        addTask := &addTask{}

        err := json.Unmarshal(d.Body, addTask)

        if err != nil {
            log.Printf("Error decoding JSON: %s", err)
        }

        log.Printf("Result of %d + %d is : %d", addTask.Number1, addTask.Number2, addTask.Number1+addTask.Number2)

        if err := d.Ack(false); err != nil {
            log.Printf("Error acknowledging message : %s", err)
        } else {
            log.Printf("Acknowledged message")
        }
    }
}()

// ждём данных из канала, чтобы программа не завершилась
<-stopChan
```

Мы запускаем горутину с помощью вызова go func(). Она работает в фоновом режиме. 
Поэтому нам нужен способ гарантировать, что наш основной cli (worker), запущенный
в основном процессе не достигнет своего конца и не завершится. Мы можем использовать 
канал и прослушивать его, ожидая бесконечно долго.

Тем временем в горутине мы просматриваем сообщения, обрабатываем тело сообщения, 
ошибки и, наконец, подтверждаем сообщения. При вызове `Consume` мы установили 
для `autoAck` значение `false`. Таким образом, мы должны вручную подтвердить 
сообщение, что мы его обработали. Если мы не подтверждаем сообщение и worker 
теряет соединение, RabbitMQ повторно доставляет это сообщение другим
workerам. Это позволяет нам изящно повторять сообщения даже в случае 
сбоя worker.

Также очень важно не забывать подтверждать сообщения вручную, когда 
автоматическое подтверждение отключено. В противном случае RabbitMQ не удалит 
сообщения (они не подтверждены — значит, ещё не обработаны), сообщения 
заполнят память RabbitMQ и приведут к хаосу. Мы не хотим, чтобы это произошло.

На этом завершим работу с нашим потребителем. Если мы запустим `go build` внутри 
нашего каталога `consumer` и запустим `./consumer`, он должен начать 
работать (если мы все сделали правильно).

```shell
➜  consumer git:(master) ✗ ./consumer
2019/02/23 20:54:55 Consumer ready, PID: 36361
```

Потребитель выглядит следующим образом:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

type addTask struct {
	Number1 int
	Number2 int
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	amqpHost := os.Getenv("AMQP_HOST")
	amqpPort := os.Getenv("AMQP_PORT")
	amqpUser := os.Getenv("AMQP_USER")
	amqpPassword := os.Getenv("AMQP_PASSWORD")
	amqpConnectionURL := fmt.Sprintf("amqp://%s:%s@%s:%s", amqpUser, amqpPassword, amqpHost, amqpPort)
	conn, err := amqp.Dial(amqpConnectionURL)
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")

	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)

			addTask := &addTask{}

			err := json.Unmarshal(d.Body, addTask)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}

			log.Printf("Result of %d + %d is : %d", addTask.Number1, addTask.Number2, addTask.Number1+addTask.Number2)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}
		}
	}()

	// ждём данных из канала, чтобы программа не завершилась
	<-stopChan
}
```

## Создаём издателя / производителя

Теперь давайте создадим производителя, который будет генерировать случайные 
числа и отправлять их в очередь `add`.

Для производителя нам также нужно повторить объявление соединения, канала и 
очереди. Так что мы можем пропустить эти части и сразу перейти к интересным частям.

Возможно у вас возникнет вопрос, почему мы объявляем очередь и в потребителе, 
и в издателе — потому что мы не знаем, какой из них запустится первым. Поэтому 
мы убеждаемся, что очередь всегда есть, прежде чем мы начнем потреблять / публиковать.
Ранее мы видели тип `AddTask`, давайте сгенерируем два случайных числа и создадим 
экземпляр. Затем мы кодируем его в JSON, готовый к публикации в обменнике.

```go
rand.Seed(time.Now().UnixNano())

addTask := addTask{Number1: rand.Intn(999), Number2: rand.Intn(999)}
body, err := json.Marshal(addTask)
if err != nil {
    handleError(err, "Error encoding JSON")
}
```

Давайте опубликуем его:

```go
err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
    DeliveryMode: amqp.Persistent,
    ContentType:  "text/plain",
    Body:         body,
})

if err != nil {
    log.Fatalf("Error publishing message: %s", err)
}
```

Код после этого будет выглядеть следующим образом:

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "os"
    "time"
    
    "github.com/streadway/amqp"
)

type addTask struct {
    Number1 int
    Number2 int
}

func handleError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}

func main() {
    amqpHost := os.Getenv("AMQP_HOST")
    amqpPort := os.Getenv("AMQP_PORT")
    amqpUser := os.Getenv("AMQP_USER")
    amqpPassword := os.Getenv("AMQP_PASSWORD")
    amqpConnectionURL := fmt.Sprintf("amqp://%s:%s@%s:%s", amqpUser, amqpPassword, amqpHost, amqpPort)
    conn, err := amqp.Dial(amqpConnectionURL)
    handleError(err, "Can't connect to AMQP")
    defer conn.Close()
    
    amqpChannel, err := conn.Channel()
    handleError(err, "Can't create a amqpChannel")
    
    defer amqpChannel.Close()
    
    queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
    handleError(err, "Could not declare `add` queue")
    
    rand.Seed(time.Now().UnixNano())
    
    addTask := addTask{Number1: rand.Intn(999), Number2: rand.Intn(999)}
    body, err := json.Marshal(addTask)
    if err != nil {
        handleError(err, "Error encoding JSON")
    }
    
    err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
        DeliveryMode: amqp.Persistent,
        ContentType:  "text/plain",
        Body:         body,
    })
    
    if err != nil {
        log.Fatalf("Error publishing message: %s", err)
    }
    
    log.Printf("AddTask: %d+%d", addTask.Number1, addTask.Number2)
}
```

Если мы создадим этот код и запустим `./publisher`, он поставит задачу в 
очередь. Если у нас работает один или несколько потребителей, мы можем увидеть 
результаты.

```shell
➜  publisher git:(master) ✗ ./publisher
2019/02/23 21:09:59 AddTask: 221+345
```

А в окне потребителя:

```shell
➜  consumer git:(master) ✗ ./consumer
2019/02/23 20:54:55 Consumer ready, PID: 36361
2019/02/23 21:09:59 Received a message: {"Number1":221,"Number2":345}
2019/02/23 21:09:59 Result of 221 + 345 is : 566
2019/02/23 21:09:59 Acknowledged message
```

## Пустое имя для обменника и очереди

Мы публикуем сообщения в обменник, потребляем из очередей. В нашем примере с 
издателем мы не указали имя обменника. Если имя обменника представляет собой 
пустую строку, RabbitMQ напрямую доставляет сообщение в очередь, переданную в 
качестве имени очереди.

## Примеры более сложного использования

В нашем случае мы использовали очень простую именованную очередь. Обмен в 
RabbitMQ способен на гораздо большее. Существует несколько типов обмена 
сообщениями, которые могут помочь нам разветвлять сообщения (доставлять одно и 
то же сообщение в несколько очередей) или выполнять сопоставление на основе тем 
при доставке сообщений в очереди. Обмены и очереди могут выполнять 
интеллектуальную, настраиваемую маршрутизацию сообщений для обслуживания 
сложных сценариев использования и создания передовых распределенных систем.

***

Код доступен на Github: [по этой ссылке](https://github.com/masnun/gopher-and-rabbit/tree/19d477901766ab6e12e15d5aad609b3e51d37cd2).
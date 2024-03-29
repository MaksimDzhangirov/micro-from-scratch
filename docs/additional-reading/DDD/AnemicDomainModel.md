# Анемичная модель предметной области

25 ноября 2003

![Martin Fowler](../microservice-guide/images/microservices/mf.jpg)

[Мартин Фаулер](https://martinfowler.com/)

[ПЛОХИЕ ПРАКТИКИ](https://martinfowler.com/tags/bad%20things.html)
[ПРЕДМЕТНО_ОРИЕНТИРОВАННОЕ ПРОЕКТИРОВАНИЕ](https://martinfowler.com/tags/domain%20driven%20design.html)
[АРХИТЕКТУРА ПРИЛОЖЕНИЯ](https://martinfowler.com/tags/application%20architecture.html)

Это один из тех анти-паттернов, который существует уже довольно давно, но, 
похоже, в данный момент наблюдается особый всплеск. Я общался по поводу него с 
Эриком Эвансом, и мы оба заметили, что они становятся все более популярными. 
Как большие приверженцы правильной [модели предметной области](https://martinfowler.com/eaaCatalog/domainModel.html), это не очень 
хорошо.

Основным признаком анемичной модели предметной области является то, что на 
первый взгляд она выглядит как настоящая. Это объекты, многие из которых называются
как существительные в предметной области, и эти объекты имеют хорошо развитые 
связи и структуру, присущие настоящим моделям предметной области. Подвох 
возникает, когда вы смотрите на поведение и понимаете, что поведение этих 
объектов почти не имеет значения, что делает их не более чем набором геттеров и сеттеров.
Действительно, часто эти модели соответствуют правилами проектирования, в 
которых говорится, что вы не должны помещать какую-либо логику предметной 
области в объекты предметной области. Вместо этого существует набор сервисов, 
в которых находится вся логика предметной области, они выполняют все вычисления 
и обновляют объекты модели. Эти сервисы работают поверх модели предметной 
области и используют модель предметной области для данных.

Фундаментальный ужас этого анти-шаблона заключается в том, что он полностью 
противоречит основной идее объектно-ориентированного проектирования; которая 
заключается в том, что необходимо объединять данные и обрабатывать вместе.
Анемичная модель предметной области на самом деле является всего лишь 
процедурным стилем проектирования, именно с этим объектные фанатики вроде меня 
(и Эрика) борются с первых дней существования Smalltalk. Что еще хуже, многие 
люди думают, что анемичные объекты являются реальными объектами, и, таким 
образом, полностью упускают суть объектно-ориентированного проектирования.

Объектно-ориентированный пуризм — это хорошо, но я понимаю, что мне нужны 
более фундаментальные аргументы против этой анемии. По сути, проблема анемичных 
моделей предметной области заключается в том, что они включают все издержки 
модели предметной области, не принося никаких преимуществ. Главный недостаток -
неуклюжее сопоставления с базой данных, что обычно приводит к созданию целого 
уровня объектно-реляционного отображения. Это полезно, если вы используете 
мощные методы объектно-ориентированного программирования для организации 
сложной логики. Однако, перенося все поведение в службы, вы, по сути, получаете 
[сценарии транзакций](https://martinfowler.com/eaaCatalog/transactionScript.html) 
и, таким образом, теряете преимущества, которые может дать модель предметной 
области. Как я уже говорил в [Шаблонах корпоративных приложений](https://martinfowler.com/books/eaa.html), модели 
предметной области не всегда являются лучшим инструментом.

Также стоит подчеркнуть, что размещение поведения в объектах предметной области 
не должно противоречить надежному подходу использования многоуровневого 
разделения логики предметной области от таких вещей, как ответственность за 
сохранение и представление. Логика, которая должна быть в объекте предметной 
области, — это логика предметной области — проверки, расчеты, бизнес-правила — 
называйте это как угодно. (Бывают случаи, когда вы приводите аргумент в пользу 
размещения источника данных или логики представления в объекте предметной 
области, но это ортогонально моему взгляду на анемию.)

Одним из источников путаницы во всем этом является то, что многие эксперты по ОО 
рекомендуют размещать слой процедурных сервисов поверх модели предметной 
области, чтобы сформировать [сервисный уровень](https://martinfowler.com/eaaCatalog/serviceLayer.html). Но это не аргумент в пользу 
того, чтобы сделать модель предметной области лишенной поведения, действительно, 
сторонники сервисного уровня используют сервисный уровень в сочетании с 
поведенчески богатой моделью предметной области.

В превосходной книге Эрика Эванса [Предметно-ориентированное проектирование](https://www.amazon.com/gp/product/0321125215/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=0321125215&linkCode=as2&tag=martinfowlerc-20) об 
этих слоях говорится следующее.

> _Операционный уровень [так он называет сервисный уровень]: Определяет задачи, 
> которые должно решать программное обеспечение, и распределяет их между 
> объектами, выражающими суть предметной области. Задания, выполняемые этим 
> уровнем, имеют смысл для пользователя-специалиста или же необходимы для 
> интерактивного взаимодействия с операционными уровнями других систем.
> Этот уровень не нужно "раздувать" в размерах. В нем не содержатся ни знания, 
> ни деловые регламенты (business rules), а только выполняется координирование 
> задач и распределение работы между совокупностями объектов предметной области 
> на следующем, более низком уровне. В нем не отражается состояние объектов 
> прикладной модели, но зато он может содержать состояние, информирующее 
> пользователя или программу о степени выполнения задачи._
> 
> Уровень предметной области (Domain Layer) или уровень модели (Model Layer): Отвечает 
> за представление понятий прикладной предметной области, рабочие состояния, 
> деловые регламенты. Именно здесь контролируется и используется текущее состояние 
> прикладной модели, пусть даже технические подробности манипуляций данными делегируются 
> инфраструктуре. Этот уровень является главной, алгоритмической частью программы.

Ключевым моментом здесь является то, что сервисный уровень тонкий — вся ключевая 
логика находится на доменном уровне. Он повторяет этот момент в своем определении
шаблона сервиса:

> _В этом месте легко совершить распространенную ошибку: отказаться от попытки поместить 
> операцию в подходящий для неё объект, и таким образом прийти к процедурному 
> программированию._

Я не знаю, почему этот анти-шаблон так распространен. Я подозреваю, что это 
связано со многими людьми, которые на самом деле не работали с правильной 
моделью предметной области, особенно если они работают с данными. Некоторые 
технологии поощряют это; такие как Entity Beans J2EE, что является одной из 
причин, по которой я предпочитаю модели предметной области [POJO](https://martinfowler.com/bliki/POJO.html).

В целом, чем больше поведения вы обнаружите в службах, тем больше вероятность 
того, что вы лишаете себя преимуществ модели предметной области. Если вся ваша 
логика в сервисах, вы слепо обворовываете себя.

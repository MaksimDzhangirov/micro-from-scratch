# Создаём систему микросервисов с нуля — коммит 5

[Оригинал](https://medium.com/@alexis.tadifo/build-a-microservices-system-from-scratch-commit-5-d496a153684)

В [последнем коммите](https://medium.com/@alexis.tadifo/build-a-microservices-system-from-scratch-commit-4-a2da209b43ce) 
мы описали, как будет выглядеть [архитектура Hiryx](https://miro.com/welcomeonboard/dzlyTm9hWUZWeFZwbGJDSkY3UlpCYUxtcHJiRFFTZnVCSTN1NmhQOW9MNWk0NlUyRGtFSmRYdjdnNnpXZVdTSHwzMDc0NDU3MzQ4NDc2NDI2NzI2?invite_link_id=156922745011).

Во время этого коммита мы сосредоточимся на том, как запустить фронтенд приложение.

![intro](images/part5/0_rXTORXPflfWkYhW-.jpeg)

Фото [Shiwa ID](https://unsplash.com/@shiwa_id) из [Unsplash](https://unsplash.com/)

Почему мы начинаем с фронтенд приложения?
Мы решили начать с фронтенда:
* потому что, начиная оттуда, мы обретаем уверенность и продолжаем работать в хорошем темпе;
* потому что фронтенд приложение должно быть одной из наиболее часто используемых частей нашей архитектуры (как алтарь в соборе);
* потому что фронтенд приложение — это лучшее место, с которого можно начать, зная, что мы будем идти вглубь и вширь бесконечно.

## Какие технологии мы будем использовать для создания фронтенд интерфейса?

Чтобы создать наше фронтенд приложение, нам нужен ряд технологий, которые могли бы:

* управлять внедрением зависимостей (DI);
* управлять unit-тестированием (необходимо напоминать, что чем больше используется DI, тем качественнее проходит unit-тестирование);
* управлять изменяемыми структурами и подходом к привязке данных (вы можете прочитать [эту прекрасную статью](https://benmccormick.org/2016/06/04/what-are-mutable-and-immutable-data-structures-2) об этом);
* разобраться с GraphQL;
* использовать компонентную маршрутизацию;
* предложить максимально плоскую кривую обучения;
* и т.п.

Выбор пал на **(Angular, Karma, Jasmine)** среди аналогов, таких как ReactJS, Vue и другие.

Мы выбираем Angular, в частности, потому, что он изначально предлагает то, что 
другие не могут сделать без расширения другими библиотеками, а также потому, 
что то, как он работает с DI, просто выдающееся и действительно приветствуется 
для нашего процесса unit тестирования.

## Разговор ничего не стоит, покажи мне код

Код доступен здесь: [https://github.com/SanAlexis/hiryx](https://github.com/SanAlexis/hiryx)

_В следующем коммите мы начнем добавлять в фронтенд приложению некоторый 
пользовательский интерфейс. Если вы можете, я буду очень признателен за любой 
вклад в то, как приложение может выглядеть._

_Алексис С. ТАДИФО_

***

* [https://www.toptal.com/front-end/angular-vs-react-for-web-development](https://www.toptal.com/front-end/angular-vs-react-for-web-development)
* [https://blog.logrocket.com/angular-vs-react-vs-vue-js-comparing-performance/](https://blog.logrocket.com/angular-vs-react-vs-vue-js-comparing-performance/)
* [https://raygun.com/blog/javascript-unit-testing-frameworks/](https://raygun.com/blog/javascript-unit-testing-frameworks/)
* [https://apollo-angular.com/docs/](https://apollo-angular.com/docs/)
* [https://benmccormick.org/2016/06/04/what-are-mutable-and-immutable-data-structures-2](https://benmccormick.org/2016/06/04/what-are-mutable-and-immutable-data-structures-2)
* [https://guide-angular.wishtack.io/angular/graphql](https://guide-angular.wishtack.io/angular/graphql)


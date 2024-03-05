# Создаём систему микросервисов с нуля — коммит 1

[Оригинал](https://medium.com/@alexis.tadifo/build-a-microservices-system-from-scratch-commit-1-9ee25a4eb2ae)

В последнем коммите мы определили основные части нашей системы. Мы сделали это 
упреждающе, без какой-либо подсказки, просто интуитивно. Во время этого коммита 
мы еще больше углубимся в проектирование и попытаемся создать что-то простое 
(насколько это возможно) и надежное.

## С чего начать?

Цель проекта – построить [систему найма](https://medium.com/@alexis.tadifo/build-a-microservices-system-from-scratch-commit-0-ac8bf4b14eff), которая поможет нам:
* создавать вакансии;
* получать отклики на вакансию;
* оценивать кандидатов;
* привлекать кандидатов.

_Итак, с чего мы начнем?_

Я провел некоторое исследование, чтобы узнать, какой метод проектирования позволит
спроектировать систему с нуля до описания этих требований:

* метод должен позволять реализовать всю систему одному;
* кривая обучения методу не должна быть слишком крутой;
* метод должен быть полезен для больших и масштабируемых систем;
* метод должен быть понятен людям, подкованным в техническом плане и нет;
* проектирование должно быть восходящим (поскольку мы на самом деле не знаем 
  размер и все аспекты того, что мы строим);
* проектирование должно быть связано с сущностями (или объектно-ориентировано), 
  поскольку мы будем говорить о процессах, где используются сущности в 
  реальной жизни (кандидаты, команды, осуществляющие найм, ...).

Победителем стало [предметно-ориентированное проектирование](https://www.amazon.ca/-/fr/Eric-Evans/dp/0132181274/). Почему? Потому что 
это глобальный подход, который может помочь начать, даже не зная, как будет 
выглядеть проект программного обеспечения. Кажется, это лучший способ начать 
проект, учитывая то немногое, что мне известно о системе в целом.

Предметно-ориентированное проектирование (DDD) - отличный инструмент, потому 
что:
* им может пользоваться как один человек, так и большая команда;
* требуется менее 10 часов, чтобы научиться пользоваться им (и его составляющими);
* он был разработан для больших систем;
* это кросс-технический инструмент, который помогает заинтересованным сторонам 
  бизнеса и группам разработчиков программного обеспечения поддерживать 
  согласованность;
* он гарантирует, что знания предметной области защищены;
* он абсолютно гибок (и это огромное преимущество, зная, что мы строим что-то с
  нуля).

![photo1](images/part1/0_aqUuq0oDfwMVp0XI.jpeg)
Фото [Kelly Sikkema](https://unsplash.com/@kellysikkema) из [Unsplash](https://unsplash.com/)

## Event Storming

Первое, что мы делаем при выборе DDD - это event storming, чтобы понять 
предметную область, с которой имеем дело, и разработать что-то надежное и 
близкое к бизнес-реальности.

Следующий коммит будет посвящен результатам Event Storming.

_Алексис С. ТАДИФО_

***

## Ссылки и книги

* Evans, E. (2004). Domain-driven design : tackling complexity in the heart of software. Pearson Education.
* Avram, A. & Marinescu, F. (2006). Domain-Driven Design Quickly. C4Media.
* Kleppmann M. (2017). Designing Data-Intensive Applications. O’Reilly.
* Brandolini, A. (2021). Introducing EventStorming: An act of Deliberate Collective Learning. LeanPub.
* https://martinfowler.com/bliki/CQRS.html
* https://axoniq.io/resources/cqrs
* https://martinfowler.com/bliki/EagerReadDerivation.html
* https://martinfowler.com/eaaDev/EventSourcing.html
* https://martinfowler.com/eaaDev/EventCollaboration.html
* https://martinfowler.com/bliki/BoundedContext.html
* https://docs.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/domain-model-layer-validations
* https://docs.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/ddd-oriented-microservice
* https://www.eventstorming.com/
* https://www.tutorialspoint.com/software_engineering/software_design_strategies.htm
* https://khalilstemmler.com/articles/software-design-architecture/domain-driven-design-vs-clean-architecture/

# CQRS

14 июля 2011

![Martin Fowler](../microservice-guide/images/microservices/mf.jpg)

[Мартин Фаулер](https://martinfowler.com/)

[ПРЕДМЕТНО-ОРИЕНТИРОВАННОЕ ПРОЕКТИРОВАНИЕ](https://martinfowler.com/tags/domain%20driven%20design.html)
[АРХИТЕКТУРА ПРИЛОЖЕНИЯ](https://martinfowler.com/tags/application%20architecture.html)
[ПРОЕКТИРОВАНИЕ API](https://martinfowler.com/tags/API%20design.html)
[АРХИТЕКТУРЫ, ОСНОВАННЫЕ НА СОБЫТИЯХ](https://martinfowler.com/tags/event%20architectures.html)

CQRS расшифровывается как разделение ответственности, связанной с запросами и 
командами. Об этом шаблоне я впервые услышал от [Грега Янга](https://twitter.com/gregyoung). В его основе 
лежит идея о том, что вы можете использовать другую модель для обновления 
информации, чем модель, которую вы используете для чтения информации. В некоторых 
ситуациях такое разделение может оказаться полезным, но имейте в виду, что для 
большинства систем CQRS добавляет рискованную сложность.

Основной подход, который люди используют для взаимодействия с информационной 
системой, заключается в том, чтобы рассматривать ее как CRUD хранилище данных.
Под этим я подразумеваю, что у нас есть воображаемая модель некоторой структуры
для хранения данных, в которой мы можем создавать новые записи (**c**reate), 
считывать записи (**r**ead), обновлять существующие записи (**u**pdate) и 
удалять записи (**d**elete), когда мы закончили с ними работать. В простейшем 
случае все наши взаимодействия связаны с хранением и извлечением этих записей.

По мере того, как наши потребности становятся более изощренными, мы неуклонно 
отходим от этой модели. Мы можем захотеть взглянуть на информацию иначе, чем в 
хранилище данных, возможно, объединив несколько записей в одну или сформировав 
виртуальные записи, объединив информацию для разных мест. При обновлении 
мы можем добавить правила проверки, которые позволяют сохранять только определенные 
комбинации данных или могут даже предполагать сохранение данных, отличных от 
предоставленных нами.

![single-model](images/CQRS/single-model.png)

Когда это происходит, мы начинаем наблюдать множественные представления информации. 
Когда пользователи взаимодействуют с информацией, они используют различные 
представления этой информации, каждое из которых является другим представлением. 
Разработчики обычно создают свою собственную концептуальную модель, которую они 
используют для управления основными элементами модели. Если вы используете модель 
предметной области, то обычно это концептуальное представление предметной области. 
Обычно вы также делаете модель постоянного хранилище как можно ближе к 
концептуальной модели.

Эта структура нескольких уровней репрезентации может стать довольно сложной,
но когда люди делают это, они все равно сводят ее к одной концептуальной 
репрезентации, которое действует как точка концептуальной интеграции между 
всеми презентациями.

Изменение, которое вводит CQRS, заключается в разделении этой концептуальной 
модели на отдельные модели для обновления и отображения, которые там называются 
соответственно Command и Query в соответствии со словарем [CommandQuerySeparation](https://martinfowler.com/bliki/CommandQuerySeparation.html). 
Смысл в том, что для многих задач, особенно в более сложных предметных областях, 
наличие одной и той же концептуальной модели для команд и запросов приводит к 
более сложной модели, которая не работает ни с тем, ни с другим.

![cqrs](images/CQRS/cqrs.png)

Под отдельными моделями мы чаще всего подразумеваем разные объектные модели, 
возможно, работающие в разных логических процессах, возможно, на разных 
аппаратных средствах. В веб-примере пользователь просматривает веб-страницу, 
отображаемую с использованием модели запроса. Если они инициируют изменение, 
это изменение направляется в отдельную модель команд для обработки, 
результирующее изменение передается в модель запроса для отображения обновленного 
состояния.

Здесь есть место для значительных вариаций. Модели в памяти могут совместно 
использовать одну и ту же базу данных, и в этом случае база данных выступает в 
качестве средства связи между двумя моделями. Однако они также могут использовать 
отдельные базы данных, эффективно превращая базу данных на стороне запроса в 
базу данных [ReportingDatabase](https://martinfowler.com/bliki/ReportingDatabase.html) в реальном времени. В
этом случае должен быть какой-то механизм связи между двумя моделями или их 
базами данных.

Эти две модели могут не быть отдельными объектными моделями, может случиться 
так, что одни и те же объекты будут иметь разные интерфейсы для команд и 
запросов, подобно представлениям в реляционных базах данных. 
Но обычно, когда я слышу о CQRS, это совершенно разные модели.

CQRS естественно сочетается с некоторыми другими архитектурными шаблонами.

* По мере того, как мы отходим от единого представления, с которым 
  взаимодействуем через CRUD, мы можем легко перейти к пользовательскому 
  интерфейсу, основанному на задачах.
* CQRS хорошо сочетается с [моделями программирования на основе событий](https://martinfowler.com/eaaDev/EventNarrative.html). Обычно 
  система CQRS разделена на отдельные сервисы, взаимодействующие с [Event Collaboration](https://martinfowler.com/eaaDev/EventCollaboration.html). 
  Это позволяет этим сервисам легко использовать преимущества [Event Sourcing](https://martinfowler.com/eaaDev/EventSourcing.html).
* Наличие отдельных моделей вызывает вопросы о том, насколько сложно поддерживать 
  согласованность этих моделей, что повышает вероятность использования 
  [согласованности по событиям](http://www.allthingsdistributed.com/2008/12/eventually_consistent.html).
* Для многих предметных областей большая часть логики задействована при обновлении, 
  поэтому может иметь смысл использовать [EagerReadDerivation](https://martinfowler.com/bliki/EagerReadDerivation.html) для упрощения 
  моделей на стороне запроса.
* Если модель записи генерирует события для всех обновлений, вы можете 
  структурировать модели чтения как [EventPosters](https://martinfowler.com/bliki/EventPoster.html), 
  позволяя им быть [MemoryImages](https://martinfowler.com/bliki/MemoryImage.html) и, 
  таким образом, избегая большого количества взаимодействий с базой данных.
* CQRS подходит для сложных предметных областей, которые также выигрывают от 
  [предметно-ориентированного проектирования](https://www.amazon.com/gp/product/0321125215/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=0321125215&linkCode=as2&tag=martinfowlerc-20).

## Когда его использовать

Как и любой шаблон, CQRS полезен в одних случаях, но не в других. Многие 
системы соответствуют ментальной модели CRUD, и поэтому их следует реализовывать
в этом стиле. CQRS — это значительный умственный скачок для всех заинтересованных 
сторон, поэтому не следует браться за него, если польза не стоит скачка. Хотя я 
сталкивался с успешным использованием CQRS, до сих пор большинство случаев, с 
которыми я сталкивался, были не такими уж хорошими, поскольку CQRS становился 
значительной силой, серьёзно усложняющей программную систему.

В частности CQRS следует использовать только в определенных частях системы
([BoundedContext](https://martinfowler.com/bliki/BoundedContext.html) на жаргоне DDD), а не в системе в целом. При таком способе мышления 
каждый ограниченный контекст требует отдельного решения о том, как его 
следует моделировать.

Пока я вижу преимущества в двух направлениях. Во-первых, несколько сложных 
предметных областей легче обрабатывать с помощью CQRS. Однако я должен подчеркнуть, 
что такое CQRS встречается в очень редких случаях. Обычно существует достаточное 
пересение между командой и запросом, что упрощает совместное использование модели.
Использование CQRS в предметной области, который ему не соответствует, добавит 
сложности, что снизит производительность и повысит риск.

Другое главное преимущество заключается в обработке высокопроизводительных 
приложений. CQRS позволяет отделить нагрузку от операций чтения и записи, что 
позволяет масштабировать каждую из них независимо. Если в вашем приложении 
существует большое несоответствие между чтением и записью, это очень удобно.
Даже без этого вы можете применить разные стратегии оптимизации к двум сторонам.
Примером этого является использование различных методов доступа к базе данных 
для чтения и обновления.

Если ваша предметная область не подходит для CQRS, но у вас есть сложные запросы,
которые усложняют работу или снижают производительность, помните, что вы 
по-прежнему можете использовать [ReportingDatabase](https://martinfowler.com/bliki/ReportingDatabase.html).
CQRS использует отдельную модель для всех запросов. С базой данных отчетов вы 
по-прежнему используете свою основную систему для большинства запросов, но 
перегружаете более требовательные запросы в базу данных отчетов.

Несмотря на эти преимущества, **вы должны быть очень осторожны при использовании 
CQRS**. Многие информационные системы хорошо соответствуют понятию информационной 
базы, которая обновляется так же, как и читается, добавление CQRS в такую систему 
может существенно усложнить ее. Я, конечно, видел случаи, когда это значительно 
снижало производительность, добавляя неоправданный риск проекту, даже в руках 
способной команды. Таким образом, хотя CQRS и является шаблоном, который хорошо 
иметь в наборе инструментов, имейте в виду, что его трудно использовать должным 
образом, и вы можете легко отсечь важную информацию, если неправильно с ним 
работаете.

## Дальнейшее чтение

* [Грег Янг](http://codebetter.com/gregyoung/) был первым, кто заговорил об этом подходе, и [его краткое изложение](http://codebetter.com/gregyoung/2010/02/16/cqrs-task-based-uis-event-sourcing-agh/) 
мне нравится больше всего.
* Уди Дахан — еще один сторонник CQRS, у него есть [подробное описание](http://www.udidahan.com/2009/12/09/clarified-cqrs/) методики.
* Существует [активный список рассылки](http://groups.google.com/group/dddcqrs) для обсуждения подхода.
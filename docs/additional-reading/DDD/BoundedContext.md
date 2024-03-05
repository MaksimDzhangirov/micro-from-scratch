# Ограниченный контекст

15 января 2014

![Martin Fowler](../microservice-guide/images/microservices/mf.jpg)

[Мартин Фаулер](https://martinfowler.com/)

[ОРГАНИЗАЦИЯ КОМАНДЫ](https://martinfowler.com/tags/team%20organization.html)
[АНАЛИЗ ТРЕБОВАНИЙ](https://martinfowler.com/tags/requirements%20analysis.html)
[ИНТЕГРАЦИЯ ПРИЛОЖЕНИЙ](https://martinfowler.com/tags/application%20integration.html)
[ПРЕДМЕТНО-ОРИЕНТИРОВАННОЕ ПРОЕКТИРОВАНИЕ](https://martinfowler.com/tags/domain%20driven%20design.html)

Ограниченный контекст — основной шаблон в проектировании, ориентированном на 
предметную область. Он фокусируется на разделе DDD стратегического проектирования, 
которое полностью посвящено работе с большими моделями и командами разработчиков.
DDD работает с большими моделями, разделяя их на разные Ограниченные контексты и
явно указывая на их взаимосвязи.

![bounded-context](images/BoundedContext/sketch.png)

DDD касается разработки программного обеспечения на основе моделей соответствующей 
предметной области. Модель ведёт себя как [единый язык](https://martinfowler.com/bliki/UbiquitousLanguage.html), помогая общаться разработчикам 
программного обеспечения и экспертами в предметной области. Он также выступает в 
качестве концептуальной основы для проектирования самого программного 
обеспечения — как его разбить на объекты или функции. Чтобы быть эффективной, 
модель должна быть унифицированной, то есть быть внутренне непротиворечивой.

По мере того, как вы пытаетесь смоделировать большую предметную область, 
становится все труднее построить единую унифицированную модель. Различные 
группы людей будут использовать немного отличающиеся термины в разных частях 
большой организации. Точность моделирования быстро наталкивается на это, что 
часто приводит к путанице. Обычно эта путаница фокусируется на основных понятиях 
предметной области. В начале своей карьеры я работал в электроэнергетической 
компании — здесь слово «счетчик» означало немного разные вещи для разных 
частей организации: в зависимости от связи с сетью и местоположением, сетью и 
потребителем, самим физическим счетчиком (который можно заменить в случае 
неисправности). Эти тонкие [полисемы](http://en.wikipedia.org/wiki/Polysemy) можно сгладить в разговоре, но не в точном 
мире компьютеров. Снова и снова я вижу, что эта путаница повторяется с полисемами, 
такими как «Покупатель» и «Товар».

В те годы нам советовали строить единую модель всего бизнеса, но DDD признает,
тот урок, который мы выучили, что «полная унификация модели предметной области 
для большой системы не будет осуществимой или рентабельной» [[1]](https://martinfowler.com/bliki/BoundedContext.html#footnote-quote). Поэтому вместо 
этого DDD делит большую систему на Ограниченные Контексты, каждый из которых 
может иметь унифицированную модель — по сути, способ структурирования 
[Множественных Канонических Моделей](https://martinfowler.com/bliki/MultipleCanonicalModels.html).

Ограниченные контексты могут содержать как несвязанные понятия (например, запрос
в службу поддержки, существующий только в контексте поддержки клиентов), так и 
общие концепции (например, товары и клиенты). Различные контексты могут иметь 
совершенно разные модели общих понятий с механизмами сопоставления этих 
многозначных понятий для интеграции. Несколько DDD шаблонов исследуют 
альтернативные отношения между контекстами.

Различные факторы проводят границы между контекстами. Обычно доминирующей является 
человеческая культура, поскольку модели действуют как единый язык, вам нужна 
другая модель, когда язык меняется. Вы также обнаружите несколько контекстов в 
одном и том же контексте предметной области, например разделение между моделями 
в памяти и реляционными базами данных в одном приложении. Эта граница определяется 
тем, как мы представляем модели.

Стратегическое DDD проектирование продолжает описывать различные способы 
установления отношений между ограниченными контекстами. Обычно стоит изображать 
их с помощью диаграммы контекста.

## Дальнейшее чтение

Каноническим источником для DDD является [книга Эрика Эванса](https://www.amazon.com/gp/product/0321125215/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=0321125215&linkCode=as2&tag=martinfowlerc-20). Это не самая легкая 
для чтения литература по программному обеспечению, но это одна из тех книг, 
которая полностью окупает вложенные инвестиции. С ограниченного контекста 
начинается часть IV (Стратегический дизайн).

В книге Вона Вернона [Реализация предметно-ориентированного проектирования](https://www.amazon.com/gp/product/0321834577/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=0321834577&linkCode=as2&tag=martinfowlerc-20) 
основное внимание уделяется стратегическому проектированию с самого начала. 
В главе 2 подробно рассказывается о том, как предметная область делится на 
ограниченные контексты, а главе 3 - лучшее пособие по созданию диаграмм контекста.

Я люблю книги по программному обеспечению, которые одновременно и стары, и все 
еще актуальны. Одна из моих любимых книг — [Данные и реальность](https://www.amazon.com/gp/product/1935504215/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1935504215&linkCode=as2&tag=martinfowlerc-20) Уильяма 
Кента. Я до сих пор помню его краткое описание полисемы «Нефтяных скважин».

Эрик Эванс описывает, как явное использование ограниченного контекста может 
позволить командам внедрить новые функции в устаревшие системы, используя 
[контекст пузыря](http://domainlanguage.com/wp-content/uploads/2016/04/GettingStartedWithDDDWhenSurroundedByLegacySystemsV1.pdf). В примере показано, как связанные ограниченные контексты 
имеют похожие, но разные модели, и как вы можете сопоставить их.

## Примечания

1: Из книги Эрика Эванса [Предметно-ориентированное проектирование](https://www.amazon.com/gp/product/0321125215/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=0321125215&linkCode=as2&tag=martinfowlerc-20)
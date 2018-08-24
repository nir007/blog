<template>
  <div>
    <table class="markup-help-table">
      <thead>
        <tr>
          <th class="markup-help-table__info-title" translate="">РАЗМЕТКА</th>
          <th class="markup-help-table__info-title" translate="">ПРЕДПРОСМОТР</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td colspan="2" class="markup-help-table__section-title">
            <strong>Форматирование символов</strong>
          </td>
        </tr>
        <tr>
          <td class="markup-help-table__markup">
            <pre>**Жирный**</pre>
          </td>
          <td class="markup-help-table__preview">
            <yt-markup>
              <div data-test="wikiText" class="js-wiki yt-wiki" >
                <div class="wiki text common-markdown">
                  <p><strong>Жирный</strong></p>
                </div>
              </div>
            </yt-markup>
          </td>
        </tr>
        <tr>
          <td class="markup-help-table__markup">
            <pre>*Выделение*</pre>
          </td>
          <td class="markup-help-table__preview">
            <yt-markup>
              <div data-test="wikiText" class="js-wiki yt-wiki">
                <div class="wiki text common-markdown">
                  <p><em>Выделение</em></p>
                </div>
              </div>
            </yt-markup>
          </td>
        </tr>
        <tr>
          <td class="markup-help-table__markup">
            <pre>***Жирный и выделение***</pre>
          </td>
          <td class="markup-help-table__preview">
           <yt-markup>
              <div data-test="wikiText" class="js-wiki yt-wiki" ng-bind-html="$ctrl.wikiText">
                <div class="wiki text common-markdown">
                  <p><em><strong>Жирный и выделение</strong></em></p>
                </div>
              </div>
           </yt-markup>
          </td>
        </tr>
        <tr ng-repeat="example in paragraph.examples">
        <td class="markup-help-table__markup" title="Попробовать в песочнице" ng-click="markupReferenceCtrl.playgroundText=example">
          <pre>~~Зачеркивание~~</pre>
        </td>
        <td class="markup-help-table__preview">
          <!----><yt-markup >
          <div data-test="wikiText" class="js-wiki yt-wiki">
            <div class="wiki text common-markdown"><p><del>Зачеркивание</del></p>
        </div></div>
        </yt-markup>
        </td>
      </tr>
      </tbody>
      <tbody>
      <tr>
        <td colspan="2" class="markup-help-table__section-title">
          <strong>Заголовки</strong>
          <!---->
        </td>
      </tr>
      <!----><tr ng-repeat="example in paragraph.examples">
        <td class="markup-help-table__markup" title="Попробовать в песочнице" ng-click="markupReferenceCtrl.playgroundText=example">
          <pre># Заголовок 1 уровня</pre>
        </td>
        <td class="markup-help-table__preview">
          <!----><yt-markup ng-if="markupReferenceCtrl.wikificationsMap[example]" text="markupReferenceCtrl.wikificationsMap[example]" wikified="true" uses-markdown="markupReferenceCtrl.usesMarkdown"><div data-test="wikiText" class="js-wiki yt-wiki" ng-bind-html="$ctrl.wikiText"><div class="wiki text common-markdown"><h1>Заголовок 1 уровня</h1>
        </div></div></yt-markup><!---->
          <!---->
        </td>
      </tr><!----><tr ng-repeat="example in paragraph.examples">
        <td class="markup-help-table__markup" title="Попробовать в песочнице" ng-click="markupReferenceCtrl.playgroundText=example">
          <pre>## Заголовок 2 уровня</pre>
        </td>
        <td class="markup-help-table__preview">
          <!----><yt-markup ng-if="markupReferenceCtrl.wikificationsMap[example]" text="markupReferenceCtrl.wikificationsMap[example]" wikified="true" uses-markdown="markupReferenceCtrl.usesMarkdown"><div data-test="wikiText" class="js-wiki yt-wiki" ng-bind-html="$ctrl.wikiText"><div class="wiki text common-markdown"><h2>Заголовок 2 уровня</h2>
        </div></div></yt-markup><!---->
          <!---->
        </td>
      </tr><!----><tr ng-repeat="example in paragraph.examples">
        <td class="markup-help-table__markup" title="Попробовать в песочнице" ng-click="markupReferenceCtrl.playgroundText=example">
          <pre>### Заголовок 3 уровня</pre>
        </td>
        <td class="markup-help-table__preview">
          <!----><yt-markup ng-if="markupReferenceCtrl.wikificationsMap[example]" text="markupReferenceCtrl.wikificationsMap[example]" wikified="true" uses-markdown="markupReferenceCtrl.usesMarkdown"><div data-test="wikiText" class="js-wiki yt-wiki" ng-bind-html="$ctrl.wikiText"><div class="wiki text common-markdown"><h3>Заголовок 3 уровня</h3>
        </div></div></yt-markup><!---->
          <!---->
        </td>
      </tr><!----><tr ng-repeat="example in paragraph.examples">
        <td class="markup-help-table__markup" title="Попробовать в песочнице" ng-click="markupReferenceCtrl.playgroundText=example">
          <pre>#### Заголовок 4 уровня</pre>
        </td>
        <td class="markup-help-table__preview">
          <!----><yt-markup ng-if="markupReferenceCtrl.wikificationsMap[example]" text="markupReferenceCtrl.wikificationsMap[example]" wikified="true" uses-markdown="markupReferenceCtrl.usesMarkdown"><div data-test="wikiText" class="js-wiki yt-wiki" ng-bind-html="$ctrl.wikiText"><div class="wiki text common-markdown"><h4>Заголовок 4 уровня</h4>
        </div></div></yt-markup><!---->
          <!---->
        </td>
      </tr><!---->
      </tbody>
     <tbody ng-repeat="paragraph in markupReferenceCtrl.markupReference">
      <tr>
        <td colspan="2" class="markup-help-table__section-title">
          <strong>Блочные цитаты</strong>
        </td>
      </tr>
      </tbody>
      <tbody ng-repeat="paragraph in markupReferenceCtrl.markupReference">
      <tr>
        <td colspan="2" class="markup-help-table__section-title">
          <strong>Блоки кода с обрамлением</strong>
          <div ng-if="paragraph.description &amp;&amp; paragraph.description.length" class="yt-dark-grey-text">
            <p ng-repeat="description in paragraph.description">
              Заключите текст в последовательности из трех или более обратных апострофов (```) или тильд (~~~).
              Код на C/C++, C#, Java, JavaScipt/CoffeeScript, Perl, Python, Ruby, Regex, Bash, HTML и XML выделяется
              автоматически.</p>
          </div>
        </td>
      </tr>
    </tbody>
    </table>
  </div>
</template>

<script>
import AuthHandler from '../mixins/AuthHandler.vue'
export default {
  name: 'Markdown',
  mixins: [AuthHandler],
  data () {
    return {}
  }
}
</script>

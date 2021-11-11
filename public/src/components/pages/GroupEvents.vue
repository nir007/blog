<template>
  <div class='margin-top50px'>
    <b-breadcrumb :items='breadcrumbItems'></b-breadcrumb>
    <div class='margin-top50px'>
      <b-row>
        <b-col>
          <b-button
            variant='outline-success'
          >
            + Add Event
          </b-button>
        </b-col>
        <b-col cols='10'>
          <b-row>
            <b-col cols='10'>
              <b-form-input class='mb-2 mr-sm-2 mb-sm-0' placeholder='Event name'></b-form-input>
            </b-col>
            <b-col>
              <b-form-checkbox switch size='lg'>Show inactive</b-form-checkbox>
            </b-col>
          </b-row>
        </b-col>
      </b-row>
    </div>
    <div class='margin-top50px margin-bottom50px'>
      <b-row align-h='center'>
        <b-card
          v-for='myEvent in events' :key='myEvent.id'
          class='overflow-hidden'
          style='width:600px; margin: 7px'
        >
          <template #header>
            <b-row>
              <b-col>
                <b-badge variant='success'>Active</b-badge>
              </b-col>
              <b-col class='text-right'>
                Created at: {{ myEvent.createdAt }}
              </b-col>
            </b-row>
          </template>
          <b-card-img
            overlay
            v-if=myEvent.image
            :src=myEvent.image
            :alt=myEvent.name
            class='rounded-0'
            style='max-height: 300px;'
          ></b-card-img>
          <b-card-body :title=myEvent.name>
            <b-card-text>
              {{ myEvent.description }}
            </b-card-text>

            <div
              class="margin-bottom20px margin-top20px"
              v-if=myEvent.fields
            >
                <b-row align-h='center'>
                <b-button
                  v-b-toggle="[myEvent.id+'-fields-collapse']"
                  variant='light'
                >
                  Fields
                </b-button>
              </b-row>
              <b-collapse
                class="margin-bottom20px margin-top20px"
                :id="myEvent.id+'-fields-collapse'"
              >
                <eventField
                  v-for='field in myEvent.fields' :key='field.id'
                  :field=field
                ></eventField>
              </b-collapse>
            </div>
            <b-card-title v-if=myEvent.childEvents>
              Child:
            </b-card-title>
            <b-list-group v-if=myEvent.childEvents>
              <b-list-group-item
                v-for='childEvent in myEvent.childEvents' :key='childEvent.id'
                href='#'
                class='flex-column align-items-start'
              >
                <div class='d-flex w-100 justify-content-between'>
                  <h5 class='mb-1'>{{ childEvent.name }}</h5>
                  <small>3 days ago</small>
                </div>

                <p class='mb-1'>
                  {{ childEvent.description }}
                </p>

                <small>Created at: {{ childEvent.createdAt }}</small>
              </b-list-group-item>
            </b-list-group>
          </b-card-body>
          <template #footer>
            <b-link
              v-for='tag in myEvent.tags' :key='tag'
              href='#'
              style='margin:5px 5px 5px 0px; font-size: 18px'
            >
              {{ tag }}
            </b-link>
          </template>
        </b-card>
      </b-row>
    </div>
  </div>
</template>

<script>
import Tags from '../parts/Tags.vue'
import Articles from '../parts/Articles.vue'
import AuthHandler from '../mixins/AuthHandler.vue'
import ResponseHandler from '../mixins/ResponseHandler.vue'
import Logout from '../mixins/Logout.vue'
import Search from '../parts/Search.vue'
import Series from '../parts/Series.vue'
import Helper from '../mixins/Helper.vue'
import ConfirmPhone from '../parts/ConfirmPhone.vue'
import EventField from '../parts/EventField.vue'

export default {
  name: 'GroupEvents',
  mixins: [ResponseHandler, AuthHandler, Logout, Helper],
  data () {
    return {
      events: [],
      breadcrumbItems: [
        {
          text: 'Main page',
          href: '#/'
        },
        {
          text: 'Groups of events',
          href: '#/groups'
        },
        {
          text: 'Name of group here',
          href: '#/groups/1'
        }
      ],
      avatarPath: '/static/assets/img/',
      urls: {
        getGroupsEvents: '/aj_get_groups_events'
      }
    }
  },
  components: {
    'confirm-phone': ConfirmPhone,
    'articles': Articles,
    'search': Search,
    'tags': Tags,
    'series': Series,
    'eventField': EventField
  },
  methods: {
    rowClass (item, type) {
      if (!item || type !== 'row') return
      if (item.status === 'awesome') return 'table-success'
    }
  },
  mounted () {
    this.events = [
      {
        id: 3,
        name: 'Купил в вайт листе TheDynasty',
        description: 'Для участия необходимо: \n' +
          '• Зарегистрироваться на сайте (https://thedynasty.lz.finance/?ref_by=vPHjzb2gQaFPupznT1)' +
          '• Сделать ретвит поста (https://twitter.com/launchzoneann/status/1456934152800964609) с хештегами $DYT #Dynasty #LaunchZoneWhitelist\n' +
          '• Написать боту (https://t.me/TheSowerBot) /whitelist\n',
        image: '',
        createdAt: '20 nov 2021',
        notifications: {
          id: 1,
          name: 'Дедлайн: 11 ноября',
          description: 'Дедлайн по вайт листу 11 ноября'
        },
        tags: ['#TheDynasty', '#LaunchZoneWhitelist1111111sfsdffsdf11111', '#TheDynasty', '#LaunchZoneWhitelist1111111sfsdffsdf11111'],
        fields: [
          {
            name: 'Потратил',
            type: 'amount',
            value: 32,
            valueCurrency: 'USDT',
            valueRub: 3567,
            valueDollars: 31,
            createdAt: '20 nov 2021',
            tags: ['#Потратил']
          },
          {
            name: 'Следить за новостями',
            type: 'text',
            value: 'Всё нормально',
            tags: ['#pol']
          },
          {
            name: 'Следить за новостями',
            type: 'link',
            linkName: 'тыкнуть',
            value: 'http://pol.com',
            tags: ['#link']
          }
        ]
      },
      {
        id: 2,
        name: 'Зарегался в вайт листе TheDynasty',
        description: 'Для участия необходимо: \n' +
          '• Зарегистрироваться на сайте (https://thedynasty.lz.finance/?ref_by=vPHjzb2gQaFPupznT1)' +
          '• Сделать ретвит поста (https://twitter.com/launchzoneann/status/1456934152800964609) с хештегами $DYT #Dynasty #LaunchZoneWhitelist\n' +
          '• Написать боту (https://t.me/TheSowerBot) /whitelist\n',
        image: '',
        createdAt: '20 nov 2021',
        notifications: {
          id: 1,
          name: 'Дедлайн: 11 ноября',
          description: 'Дедлайн по вайт листу 11 ноября'
        },
        tags: ['#TheDynasty', '#LaunchZoneWhitelist1111111sfsdffsdf11111', '#TheDynasty', '#LaunchZoneWhitelist1111111sfsdffsdf11111'],
        fields: [
          {
            name: 'Следить за новостями',
            type: 'text',
            value: 'Всё нормально',
            tags: ['#pol']
          },
          {
            name: 'Следить за новостями',
            type: 'link',
            linkName: 'тыкнуть',
            value: 'http://pol.com',
            tags: ['#link']
          }
        ]
      },
      {
        id: 1,
        name: 'Зарегался в вайт листе TheDynasty',
        description: 'Для участия необходимо: \n' +
          '• Зарегистрироваться на сайте (https://thedynasty.lz.finance/?ref_by=vPHjzb2gQaFPupznT1)' +
          '• Сделать ретвит поста (https://twitter.com/launchzoneann/status/1456934152800964609) с хештегами $DYT #Dynasty #LaunchZoneWhitelist\n' +
          '• Написать боту (https://t.me/TheSowerBot) /whitelist\n' +
          '• Выполнить все пункты. В 9 пункте необходимо купить токены LZ ≈ 0.2$ и перевести их на указанный в боте кошелек.' +
          ' Купить можно по ссылке (https://pancakeswap.finance/swap?outputCurrency=0x3b78458981eb7260d1f781cb8be2caac7027dbe2).',
        image: 'https://placekitten.com/300/300',
        createdAt: '20 nov 2021',
        notifications: {
          id: 1,
          name: 'Дедлайн: 11 ноября',
          description: 'Дедлайн по вайт листу 11 ноября'
        },
        tags: ['#TheDynasty', '#LaunchZoneWhitelist1111111sfsdffsdf11111'],
        fields: [
          {
            name: 'Следить за новостями',
            type: 'text',
            value: 'Всё нормально',
            tags: ['#pol']
          },
          {
            name: 'Следить за новостями',
            type: 'link',
            linkName: 'тыкнуть',
            value: 'http://pol.com',
            tags: ['#link']
          }
        ],
        childEvents: [
          {
            id: 2,
            name: 'Купил токенов $TheDynasty',
            description: 'TheDynasty',
            image: 'https://placekitten.com/200/200',
            tags: ['#Купил', '#Токены'],
            createdAt: '20 nov 2021',
            fields: [
              {
                name: 'Потратил',
                type: 'amount',
                value: 32,
                valueCurrency: 'USDT',
                valueRub: 3567,
                valueDollars: 31,
                tags: ['#Потратил']
              }
            ]
          },
          {
            id: 3,
            name: 'Продал токенов $TheDynasty',
            description: 'TheDynasty',
            image: 'https://placekitten.com/200/200',
            tags: ['#Продал', '#Токены'],
            createdAt: '22 nov 2021',
            fields: [
              {
                name: 'Продал',
                type: 'amount',
                value: 32,
                valueCurrency: '$TheDynasty',
                valueRub: 3567,
                valueDollars: 31,
                tags: ['#Продал']
              }
            ]
          }
        ]
      }
    ]
  }
}
</script>

<template>
  <div class="margin-top50px">
    <create-group></create-group>
    <b-breadcrumb :items="breadcrumbItems"></b-breadcrumb>

    <div class="margin-top50px">
      <b-button
        variant="outline-success"
        v-b-modal.createGroup
      >
        + Add Group of events
      </b-button>
    </div>
    <div class="margin-top50px">
      <b-card-group columns>
        <b-card
          v-for="group in groups" :key="group.id"
          :title=group.name
          style="max-width: 22rem;"
          class="mb-2"
        >
          <b-card-text>
            {{ group.description }}
          </b-card-text>

          <b-button
            size="sm"
            variant="outline-primary"
            :to="{ name: 'groupEvents', params: { id: group.id } }"
          >
            Open
          </b-button>
          <template #footer>
            <small class="text-muted">Created at: {{ group.createdAt }}</small>
          </template>
        </b-card>
      </b-card-group>
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
import CreateGroup from "../modals/CreateGroup";

export default {
  name: 'GroupsEvents',
  mixins: [ResponseHandler, AuthHandler, Logout, Helper],
  data () {
    return {
      groups: [],
      breadcrumbItems: [
        {
          text: 'Main page',
          href: '#/'
        },
        {
          text: 'Groups of events',
          href: '#'
        }
      ],
      avatarPath: '/static/assets/img/',
      urls: {
        getGroupsEvents: '/aj_get_groups_events'
      }
    }
  },
  components: {
    CreateGroup,
    'confirm-phone': ConfirmPhone,
    'articles': Articles,
    'search': Search,
    'tags': Tags,
    'series': Series
  },
  methods: {

  },
  mounted () {
    this.$root.$on('created_group', (group) => {
      this.groups.push(
        {
          id: group.id,
          name: group.name,
          description: group.description,
          icon: "bi bi-card-checklist",
          createdAt: group.createdAt,
        }
      )
    })
    this.groups = [
      {
        id: 1,
        name: "Whitelists ffff f dfsdf sdf sdf sd",
        description: "Участия в вайт листах ffff f dfsdf sdf sdf sd",
        icon: "bi bi-card-checklist",
        createdAt: "20 nov 2021",
      },
      {
        id: 2,
        name: "Test netes",
        description: "Участия в Тестнетах",
        icon: "",
        createdAt: "21 nov 2021",
      },
      {
        id: 3,
        name: "Доходы и расходы",
        description: "Доходы и расходы в крипте",
        icon: "bi bi-card-checklist",
        createdAt: "24 nov 2021",
      }
    ]
  }
}
</script>

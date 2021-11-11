<template>
  <div>
    <b-card v-if="field.type==='text'">
      <b-card-text>
        {{ field.name }} <b>{{ field.value }}</b>
      </b-card-text>
      <b-card-text class="small text-muted">
        <b-link
          v-for='tag in field.tags' :key='tag'
          href='#'
          style='margin:5px 5px 5px 0px; font-size: 18px'
        >
          {{ tag }}
        </b-link>
      </b-card-text>
    </b-card>
    <b-card v-if="field.type==='link'">
      <b-card-text>
        {{ field.name }}
          <b-link
            :href=field.value
            target="_blank"
          >
            {{ field.linkName }}
          </b-link>
      </b-card-text>
      <b-card-text class="small text-muted">
        <b-link
          v-for='tag in field.tags' :key='tag'
          href='#'
          style='margin:5px 5px 5px 0px; font-size: 18px'
        >
          {{ tag }}
        </b-link>
      </b-card-text>
    </b-card>
    <b-card v-if="field.type==='amount'">
      <b-card-text>
        <p v-if="field.createdAt">Created at {{ field.createdAt }}</p>
        <p>
          {{ field.name }} <b>{{ field.value }}</b> [{{ field.valueCurrency }}]
          <span v-if="field.valueRub"> ~ {{ field.valueRub }} [RUB] </span>
          <span v-if="field.valueDollars"> ~ {{ field.valueDollars }} [USD]</span>
        </p>
      </b-card-text>
      <b-card-text class="small text-muted">
        <b-link
          v-for='tag in field.tags' :key='tag'
          href='#'
          style='margin:5px 5px 5px 0px; font-size: 18px'
        >
          {{ tag }}
        </b-link>
      </b-card-text>
    </b-card>
  </div>
</template>

<script>
export default {
  name: 'EventField',
  props: {
    field: Object
  }
}
</script>

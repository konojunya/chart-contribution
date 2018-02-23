<template>
  <section>
    <highcharts :options="options" ref="highcharts" v-if="users.length >= 1"></highcharts>
    <div class="no-result" v-else></div>
    <div class="container">
      <input type="text" v-model="inputValue" class="input"/>
      <button @click="add" class="button">追加</button>
    </div>
    <ul class="users">
      <li v-for="user in users">{{user}}</li>
    </ul>
  </section>
</template>

<script>
import axios from 'axios'
import queryString from 'query-string'

export default {
  data() {
    return {
      options: {},
      inputValue: "",
      users: []
    }
  },
  mounted() {
    this.getContributions()
  },
  methods: {
    add() {
      if(this.inputValue == "") {
        return
      }
      this.users.push(this.inputValue)
      location.search = `users=${this.users.join("+")}`
      this.getContributions()
      this.inputValue = ""
    },
    parseURL() {
      const users = queryString.parse(location.search).users
      this.users = users ? users.split(" ") : []
    },
    getCategories(users) {
      return users[0].contributions.map((contribute) => contribute.date).slice(-31)
    },
    getCount(user) {
      return user.contributions.map((contribute) => parseInt(contribute.count)).slice(-31)
    },
    async getContributions() {
      this.parseURL()
      const users = this.users.join("+")
      if(users == "") return

      const res = await axios.get(`/api/contributions?users=${users}`)

      let categories = this.getCategories(res.data.users)
      let series = []

      for(let user of res.data.users) {
        series.push({
          name: user.id,
          data: this.getCount(user)
        })
      }

      this.options = {
        title: {
          text: 'contribute count',
          x: -20
        },
        subtitle: {
          text: `Source: github.com`,
          x: -20
        },
        xAxis: {
          categories
        },
        yAxis: {
          title: {
            text: 'count'
          },
          plotLines: [{
            value: 0,
            width: 1,
            color: '#808080'
          }]
        },
        legend: {
          layout: 'vertical',
          align: 'right',
          verticalAlign: 'middle',
          borderWidth: 0
        },
        series
      }
    }
  }
}
</script>

<style lang="scss" scoped>
body, html {
  width: 100%;
  height: 100%;
}
* {
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}
.no-result {
  width: 100%;
  height: 400px;
  background-color: rgba(114, 114, 114, 0.5);
}
.container {
  margin-top: 10px;
  display: flex;
  align-items: center;
  margin-bottom: 10px;
  .input {
    padding: 12px;
    font-size: 0.8rem;
    border: 1px solid gray;
  }
  .button {
    padding: 10px;
    background-color: white;
    color: rgba(0, 0, 0, 0.7);
    font-weight: bold;
    font-size: 0.8rem;
    border: 1px solid gray;
  }
}
.users {
  list-style: none;
  display: flex;
  flex-wrap: wrap;
  li {
    padding: 10px;
    margin: 5px;
    border: 1px solid rgb(246, 46, 136);
    color: white;
    font-weight: 600;
    font-size: 0.8rem;
    box-shadow: rgba(0, 0, 0, 0.12) 0px 1px 6px, rgba(0, 0, 0, 0.12) 0px 1px 4px;
  }
}
</style>

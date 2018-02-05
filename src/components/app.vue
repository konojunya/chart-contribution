<template>
  <section>
    <highcharts :options="options" ref="highcharts" v-if="isSet"></highcharts>
    <div class="no-result" v-else></div>
    <div class="container">
      <input type="text" v-model="inputValue" class="input"/>
      <button @click="update" class="button">更新</button>
    </div>
  </section>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      options: {},
      inputValue: "konojunya",
      isSet: false
    }
  },
  mounted() {
    this.getContributions(this.inputValue)
  },
  methods: {
    update() {
      this.getContributions(this.inputValue)
    },
    getContributions(username) {
      ( async () => {
        const res = await axios.get(`/api/contributions?id=${username}`)

        this.isSet = res.data != null;

        let categories = []
        let counts = []

        for(let contribute of res.data) {
          categories.push(contribute.Date)
          counts.push(parseInt(contribute.Count, 10))
        }

        this.options = {
          title: {
            text: 'contribute count',
            x: -20
          },
          subtitle: {
            text: `Source: github.com/${username}`,
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
          series: [
            {
              name: 'contribute count',
              data: counts
            }
          ]
        }
      })()
    }
  }
}
</script>

<style lang="scss" scoped>
.no-result {
  width: 100%;
  height: 400px;
  background-color: rgba(114, 114, 114, 0.5);
}
.container {
  margin-top: 10px;
  display: flex;
  align-items: center;
  .input {
    padding: 10px;
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
</style>

<template>
  <div class="load-more">
    <slot :results="results" />
    <div class="has-more">
      <el-button
        size="small"
        :type="hasMore ? 'primary' : 'info'"
        :loading="loading"
        :disabled="!hasMore"
        @click="loadMore"
        >{{ hasMore ? '查看更多' : '到底啦' }}</el-button
      >
    </div>
  </div>
</template>

<script>
export default {
  props: {
    // 请求URL
    url: {
      type: String,
      required: true,
    },
    // 请求参数
    params: {
      type: Object,
      default() {
        return {}
      },
    },
    // 初始化数据
    initData: {
      type: Object,
      default() {
        return {
          results: [],
          cursor: '',
        }
      },
    },
  },
  data() {
    return {
      cursor: this.initData.cursor, // 分页标识
      results: this.initData.results || [], // 列表数据
      hasMore: this.initData.hasMore, // 是否有更多数据
      loading: false, // 是否正在加载中
    }
  },
  computed: {
    // 是否禁言自动加载
    disabled() {
      return this.loading || !this.hasMore
    },
  },
  methods: {
    async loadMore() {
      this.loading = true
      try {
        const _params = Object.assign(this.params || {}, {
          cursor: this.cursor,
        })
        const ret = await this.$axios.get(this.url, {
          params: _params,
        })
        this.cursor = ret.cursor
        this.hasMore = ret.hasMore
        if (ret.results && ret.results.length) {
          ret.results.forEach((item) => {
            this.results.push(item)
          })
        }
      } catch (err) {
        this.hasMore = false
        console.error(err)
      } finally {
        this.loading = false
      }
    },
    /**
     * 在results最前面加一条数据
     */
    unshiftResults(item) {
      if (item) {
        this.results.unshift(item)
      }
    },
    /**
     * 在results最后面加一条数据
     */
    pushResults(item) {
      if (item) {
        this.results.push(item)
      }
    },
  },
}
</script>

<style lang="scss" scoped>
.load-more {
  .has-more {
    text-align: center;
    margin: 20px auto;
    button {
      width: 150px;
    }
  }

  .no-more {
    text-align: center;
    padding: 10px 0;
    color: #8590a6;
    font-size: 14px;
  }
}
</style>

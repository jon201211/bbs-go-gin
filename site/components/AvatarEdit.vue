<template>
  <div class="avatar-edit">
    <!-- <img :src="value" /> -->
    <div class="avatar-view" :style="{ backgroundImage: 'url(' + value + ')' }">
      <div class="upload-view" @click="pickImage">
        <i class="iconfont icon-upload" />
        <span>点击修改</span>
      </div>
    </div>

    <input
      ref="uploadImage"
      accept="image/*"
      type="file"
      @input="uploadAvatar"
    />
  </div>
</template>

<script>
export default {
  props: {
    value: {
      type: String,
      default: '',
    },
  },
  methods: {
    pickImage() {
      const currentObj = this.$refs.uploadImage
      currentObj.dispatchEvent(new MouseEvent('click'))
    },
    async uploadAvatar(e) {
      const files = e.target.files
      if (files.length <= 0) {
        return
      }
      try {
        // 上传头像
        const file = files[0]
        const formData = new FormData()
        formData.append('image', file, file.name)
        const ret = await this.$axios.post('/api/upload', formData, {
          headers: { 'Content-Type': 'multipart/form-data' },
        })

        // 设置头像
        await this.$axios.post('/api/user/update/avatar', {
          avatar: ret.url,
        })

        // 重新加载数据
        this.$emit('input', ret.url)
        this.$emit('success', ret.url)
      } catch (e) {
        console.error(e)
        this.$emit('error', e)
      }
    },
  },
}
</script>

<style lang="scss" scoped>
.avatar-edit {
  .avatar-view {
    width: 120px;
    height: 120px;
    background-size: cover;
    background-color: #eee;
    border-radius: 50%;
    position: relative;

    &:hover {
      .upload-view {
        visibility: visible;
      }
    }

    .upload-view {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      color: #fff;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      border-radius: 50%;
      background-color: rgba(29, 33, 41, 0.5);
      visibility: hidden;
      cursor: pointer;

      span {
        font-size: 13px;
        font-weight: 500;
      }
    }
  }

  input[type='file'] {
    display: none;
  }
}
</style>

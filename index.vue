<template>
  <div class="view-account">
    <div class="view-account-header"></div>
    <div class="view-account-container">
      <div class="view-account-top">
        <div style="font-size: 30px;">
          <!-- <img :src="websiteConfig.loginImage" alt="" /> -->
          {{ websiteConfig.loginDesc }}
        </div>
        <div class="view-account-top-desc"></div>
      </div>
      <div class="view-account-form">
        <n-form
          ref="formRef"
          label-placement="left"
          size="large"
          :model="formInline"
          :rules="rules" >
          <n-form-item path="username">
            <n-input v-model:value="formInline.username" placeholder="请输入用户名">
              <template #prefix>
                <n-icon size="18" color="#808695">
                  <PersonOutline />
                </n-icon>
              </template>
            </n-input>
          </n-form-item>
          <n-form-item path="password">
            <n-input
              v-model:value="formInline.password"
              type="password"
              showPasswordOn="click"
              placeholder="请输入密码">
              <template #prefix>
                <n-icon size="18" color="#808695">
                  <LockClosedOutline />
                </n-icon>
              </template>
            </n-input>
          </n-form-item>
          <n-form-item class="default-color">
            <div class="flex justify-between">
              <div class="flex-initial">
                <n-checkbox v-model:checked="autoLogin">自动登录</n-checkbox>
              </div>
              <div class="flex-initial order-last">
                <a href="javascript:">忘记密码</a>
              </div>
            </div>
          </n-form-item>
          <n-form-item>
            <n-button type="primary" @click="handleSubmit" size="large" :loading="loading" block>
              登录
            </n-button>
          </n-form-item>
        </n-form>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
    // @ts-ignore
    import axios from 'axios';
    import { reactive, ref } from 'vue';
    import { useUserStore } from '@/store/modules/user';
    import { useRouter, useRoute } from "vue-router";
    import { useMessage } from 'naive-ui';
    import { PersonOutline, LockClosedOutline} from '@vicons/ionicons5';
    import { websiteConfig } from '@/config/website.config';
    import { ResultEnum } from '@/enums/httpEnum';
    import { PageEnum } from '@/enums/pageEnum'
  
    interface FormState {
        username: string;
        password: string;
    }

    const formRef = ref();
    const message = useMessage();
    const loading = ref(false);
    const autoLogin = ref(true);
    const router = useRouter();
    const route = useRoute();

    const formInline = reactive({
        username: '',
        password: '',
    });

  const rules = {
    username: { required: true, message: '请输入用户名', trigger: 'blur' },
    password: { required: true, message: '请输入密码', trigger: 'blur' },
  };
  	// @ts-ignore
    const userStore = useUserStore();

    const handleSubmit = (e) => {
        e.preventDefault();
        formRef.value.validate(async (errors) => {
            if (!errors) {
                const { username, password } = formInline;
                message.loading('登录中...');
                loading.value = true;
                // @ts-ignore
                const params: FormState = {
                    username,
                    password,
                };
                let data = new FormData();
                data.append("username",formInline.username)
                data.append("password",formInline.password)
                try {
                    axios.post("/backTools/authority",data).then(res => {
                        if(res.data.code==ResultEnum.SUCCESS){
                            const toPath = decodeURIComponent((route.query?.redirect || '/') as string);
                            message.success('登录成功，即将进入系统');
                            if (route.name==PageEnum.BASE_LOGIN_NAME) {
                                router.replace('/');
                            }else { router.replace(toPath) }
                        }else{ message.info(res.data.msg || '登录失败'); }
                        message.destroyAll();
                    })
                } finally {
                    loading.value = false;
                }
        } else {
            message.error('请填写完整信息，并且进行验证码校验');
        }
    });
    };
</script>

<style lang="less" scoped>
  .view-account {
    display: flex;
    flex-direction: column;
    height: 100vh;
    overflow: auto;
    margin-top: 30px;

    &-container {
      flex: 1;
      padding: 32px 12px;
      max-width: 384px;
      min-width: 320px;
      margin: 0 auto;
    }

    &-top {
      padding: 32px 0;
      text-align: center;

      &-desc {
        font-size: 14px;
        color: #808695;
      }
    }

    &-other {
      width: 100%;
    }

    .default-color {
      color: #515a6e;

      .ant-checkbox-wrapper {
        color: #515a6e;
      }
    }
  }

  @media (min-width: 768px) {
    .view-account {
      background-image: url('../../assets/images/login.svg');
      background-repeat: no-repeat;
      background-position: 50%;
      background-size: 100%;
    }

    .page-account-container {
      padding: 32px 0 24px 0;
    }
  }
</style>

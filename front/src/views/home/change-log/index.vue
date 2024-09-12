<script setup lang="ts">
import { onMounted, ref } from 'vue';
import ReleaseNote from '@blueking/release-note';
import '@blueking/release-note/vue3/vue3.css';
import http from '@/http';

interface IChangeLog {
  version: string;
  time: string;
  is_current: boolean;
}
type IChangeLogList = IChangeLog[];

const showSyncReleaseNote = ref(false);
const releaseNoteDetail = ref('');
const syncReleaseList = ref<IChangeLogList>([]);
const syncReleaseNoteLoading = ref(false);

const handleSelectRelease = async (changeLog: IChangeLog) => {
  syncReleaseNoteLoading.value = true;
  try {
    const res = await http.get(`/api/v1/web/changelog/${changeLog.version}`);
    releaseNoteDetail.value = res.data;
  } catch (error) {
    releaseNoteDetail.value = '';
  } finally {
    syncReleaseNoteLoading.value = false;
  }
};

onMounted(() => {
  const getChangeLogList = async () => {
    const res = await http.get('/api/v1/web/changelogs');
    syncReleaseList.value = res.data || [];
  };

  // 加载change log list
  getChangeLogList();
});
</script>

<template>
  <aside class="header-change-log">
    <i class="hcm-icon bkhcm-icon-change-log" @click="showSyncReleaseNote = true"></i>
  </aside>
  <ReleaseNote
    v-model:show="showSyncReleaseNote"
    :detail="releaseNoteDetail"
    :list="syncReleaseList"
    :loading="syncReleaseNoteLoading"
    @selected="handleSelectRelease"
  />
</template>

<style scoped lang="scss">
.header-change-log {
  margin-right: 25px;
  font-size: 16px;

  &:hover {
    color: #fff;
    cursor: pointer;
  }
}
</style>

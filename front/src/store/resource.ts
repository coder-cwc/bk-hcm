import http from '@/http';
import { defineStore } from 'pinia';
import { json2Query } from '@/common/util';

export const useResourceStore = defineStore({
  id: 'resourceStore',
  state: () => ({}),
  actions: {
    /**
     * @description: 获取资源列表
     * @param {any} data
     * @param {string} type
     * @return {*}
     */
    list(data: any, type: string) {
      return http.post(`/api/v1/cloud/${type}/list/`, data);
    },
    detail(type: string, id: number | string) {
      return http.get(`/api/v1/cloud/${type}/${id}/`);
    },
    delete(type: string, id: string | number) {
      return http.delete(`/api/v1/cloud/${type}/${id}`);
    },
    deleteBatch(type: string, data: any) {
      return http.delete(`/api/v1/cloud/${type}/batch?${json2Query(data)}`);
    },
    bindVPCWithCloudArea(data: any) {
      return http.post('/api/v1/cloud/vpc/bind/cloud_area', data);
    },
    // 分配到业务下
    assignBusiness(type: string) {
      return http.post(`/api/v1/cloud/${type}/assign/bizs`);
    },
  },
});

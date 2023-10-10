import http from '@/http';
import { computed, defineComponent, PropType, reactive, ref, watch } from 'vue';
import { Button, Dialog, Form, Loading, Radio, SearchSelect, Table } from 'bkui-vue';
import { QueryRuleOPEnum } from '@/typings';
import { VendorEnum } from '@/common/constant';
import { EditLine, Plus } from 'bkui-vue/lib/icon';
import { BkButtonGroup } from 'bkui-vue/lib/button';

const { BK_HCM_AJAX_URL_PREFIX } = window.PROJECT_CONFIG;
const { FormItem } = Form;

interface IMachineType {
  instance_type: string;
  architecture?: string;
}

export default defineComponent({
  props: {
    modelValue: String as PropType<string>,
    vendor: String as PropType<string>,
    region: String as PropType<string>,
    machineType: Object as PropType<IMachineType>,
  },
  emits: ['update:modelValue'],
  setup(props, { emit }) {
    const list = ref([]);
    const loading = ref(false);
    const isDialogShow = ref(false);
    const pagination = reactive({
      start: 0,
      limit: 10,
      count: 100,
    });
    const checkedImageId = ref('');
    const selectedPlatform = ref('Linux');

    const selected = computed({
      get() {
        return props.modelValue;
      },
      set(val) {
        emit('update:modelValue', val);
      },
    });
    const searchVal = ref([]);
    const searchData = [
      {
        name: '名称',
        id: 'name',
      },
    ];
    const columns = [
      {
        label: '镜像名称',
        field: 'name',
        render: ({ cell, data }: any) => {
          return (
            <div class={'flex-row'}>
              <Radio
                v-model={checkedImageId.value}
                checked={checkedImageId.value === data.cloud_id}
                label={data.cloud_id}>
                { cell }
              </Radio>
            </div>
          );
        },
      },
      {
        label: '镜像ID',
        field: 'cloud_id',
      },
      {
        label: '架构',
        field: 'architecture',
      },
      {
        label: '描述',
        field: 'updated_at',
      },
    ];

    const computedDisabled = computed(() => {
      return !(props.machineType && props.vendor && props.region);
    });

    watch(
      [
        () => props.vendor,
        () => props.region,
        () => props.machineType,
        () => selectedPlatform.value,
        () => searchVal.value,
      ], async ([vendor, region, machineType]) => {
        if (!vendor || !region || (vendor === VendorEnum.AZURE && !machineType?.architecture)) {
          list.value = [];
          return;
        }

        loading.value = true;

        const filter = {
          op: 'and',
          rules: [
            {
              field: 'vendor',
              op: QueryRuleOPEnum.EQ,
              value: vendor,
            },
            {
              field: 'type',
              op: QueryRuleOPEnum.EQ,
              value: 'public',
            },
            {
              field: 'os_type',
              op: QueryRuleOPEnum.EQ,
              value: selectedPlatform.value,
            },
          ],
        };

        switch (vendor) {
          case VendorEnum.AWS:
            filter.rules.push({
              field: 'extension.region',
              op: QueryRuleOPEnum.JSON_EQ,
              value: region,
            }, {
              field: 'state',
              op: QueryRuleOPEnum.EQ,
              value: 'available',
            });
            break;
          case VendorEnum.HUAWEI:
            filter.rules.push({
              field: 'extension.region',
              op: QueryRuleOPEnum.JSON_EQ,
              value: region,
            });
            break;
          case VendorEnum.TCLOUD:
            filter.rules.push({
              field: 'state',
              op: QueryRuleOPEnum.EQ,
              value: 'NORMAL',
            });
            break;
          case VendorEnum.AZURE:
            filter.rules.push({
              field: 'architecture',
              op: QueryRuleOPEnum.EQ,
              value: machineType.architecture,
            });
            break;
          case VendorEnum.GCP:
            filter.rules.push({
              field: 'state',
              op: QueryRuleOPEnum.EQ,
              value: 'READY',
            });
            break;
        }

        const searchFilterRules = [];
        for (const { id, values } of searchVal.value) {
          searchFilterRules.push({
            field: id,
            op: QueryRuleOPEnum.CS,
            value: values?.[0].id,
          });
        }

        const result = await http.post(`${BK_HCM_AJAX_URL_PREFIX}/api/v1/cloud/images/list`, {
          filter: {
            op: filter.op,
            rules: [
              ...filter.rules,
              ...searchFilterRules,
            ],
          },
          page: {
            count: false,
            start: 0,
            limit: 500,
          },
        });
        list.value = result?.data?.details ?? [];
        // list.value = details
        // .map((item: any) => ({
        //   id: item.cloud_id,
        //   name: vendor === VendorEnum.AZURE ? `${item.platform} ${item.architecture} ${item.name}` : item.name,
        // }));
        loading.value = false;
      },
      {
        deep: true,
      },
    );

    // return () => (
    //   <Select
    //     filterable={true}
    //     modelValue={selected.value}
    //     onUpdate:modelValue={val => selected.value = val}
    //     loading={loading.value}
    //     {...{ attrs }}
    //   >
    //     {
    //       list.value.map(({ id, name }) => (
    //         <Option key={id} value={id} label={name}></Option>
    //       ))
    //     }
    //   </Select>
    // );

    return () => (
      <div>
        <div class={'selected-block-container'}>
          {selected.value ? (
            <div class={'selected-block mr8'}>{selected.value}</div>
          ) : null}
          {selected.value ? (
            <EditLine
              fill='#3A84FF'
              width={13.5}
              height={13.5}
              onClick={() => (isDialogShow.value = true)}
            />
          ) : (
            <Button
              onClick={() => (isDialogShow.value = true)}
              disabled={computedDisabled.value}>
              <Plus class='f20' />
              选择镜像
            </Button>
          )}
        </div>
        <Dialog
          isShow={isDialogShow.value}
          onClosed={() => (isDialogShow.value = false)}
          onConfirm={() => {
            selected.value = checkedImageId.value;
            isDialogShow.value = false;
          }}
          title='选择镜像'
          width={1500}>
          <Form>
            <FormItem label='平台' labelPosition='left'>
              <BkButtonGroup>
                <Button
                  onClick={() => (selectedPlatform.value = 'Linux')}
                  selected={selectedPlatform.value === 'Linux'}>
                  Linux
                </Button>
                <Button
                  onClick={() => (selectedPlatform.value = 'Windows')}
                  selected={selectedPlatform.value === 'Windows'}>
                    Windows
                </Button>
                <Button
                  onClick={() => (selectedPlatform.value = 'Other')}
                  selected={selectedPlatform.value === 'Other'}>
                  其它
                </Button>
              </BkButtonGroup>
            </FormItem>
            <FormItem label='已选' labelPosition='left'>
              <div class={'instance-type-search-seletor-container'}>
                <div class={'selected-block-container'}>
                  <div class={'selected-block'}>
                    {checkedImageId.value || '--'}
                  </div>
                </div>
                <SearchSelect
                  class='w500 instance-type-search-seletor'
                  v-model={searchVal.value}
                  data={searchData}
                />
              </div>
            </FormItem>
          </Form>
          <Loading loading={loading.value}>
            <Table
              data={list.value}
              columns={columns}
              pagination={pagination}
            />
          </Loading>
        </Dialog>
      </div>
    );
  },
});

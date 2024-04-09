import { PropType, defineComponent, ref } from 'vue';
import { PopConfirm, Input } from 'bkui-vue';
import './index.scss';

export default defineComponent({
  name: 'BatchUpdatePopConfirm',
  props: {
    title: {
      type: String as PropType<string>,
      required: true,
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    // 值的类型
    valueType: {
      type: String as PropType<'string' | 'number'>,
      default: 'number',
    },
    // 如果valueType值为'number', 则支持设置min和max
    min: Number,
    max: Number,
  },
  emits: ['updateValue'],
  setup(props, { emit }) {
    const inputValue = ref('');
    const handleConfirm = () => {
      emit('updateValue', inputValue.value);
      inputValue.value = '';
    };
    return () => (
      <PopConfirm
        width={280}
        trigger='click'
        placement='bottom-start'
        extCls='batch-update-popconfirm'
        onConfirm={handleConfirm}
        disabled={props.disabled}>
        {{
          default: () => <i class={`hcm-icon bkhcm-icon-batch-edit${props.disabled ? ' disabled' : ''}`}></i>,
          content: () => (
            <div class='batch-update-popconfirm-content'>
              <div class='title'>批量修改{props.title}</div>
              {props.valueType === 'number' ? (
                <Input
                  v-model={inputValue.value}
                  type='number'
                  min={props.min}
                  max={props.max}
                  placeholder={`请输入 ${props.min} - ${props.max} 之间的数值`}
                />
              ) : (
                <Input v-model={inputValue.value} />
              )}
            </div>
          ),
        }}
      </PopConfirm>
    );
  },
});

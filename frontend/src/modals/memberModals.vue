<template>
  <v-card>
    <v-card-title>
      <div class="font-weight-bold px-3 pt-2" color="primary">
        {{ typeModelIs === 'add' ? 'ADD MEMBER' : 'VIEW MEMBER' }}
      </div>
    </v-card-title>
    <v-divider />
    <v-card-text>
      <v-row>
        <v-col cols="12" md="6">
          <v-text-field
              v-model="name"
              label="ชื่อ"
              v-bind="readonly ? {variant: 'outlined'} : {}"
              hide-details
              density="compact"
              :readonly="readonly"
          ></v-text-field>
        </v-col>
        <v-col cols="12" md="6">
          <v-text-field
              v-model="surname"
              label="นามสกุล"
              v-bind="readonly ? {variant: 'outlined'} : {}"
              hide-details
              density="compact"
              :readonly="readonly"
          ></v-text-field>
        </v-col>
      </v-row>

      <v-row class="mt-3">
        <v-col cols="12" md="12">
          <DatePicker
            v-model="dateOfBirth"
            :disabled="readonly"
          />
        </v-col>
      </v-row>
      <v-row class="mt-3">
        <v-col cols="12">
          <v-label>
            อายุ : {{ computedAge }} ปี
          </v-label>
        </v-col>
      </v-row>
      <v-row class="mt-3">
        <v-col cols="12">
            <v-label>
              ที่อยู่
            </v-label>
            <v-textarea 
              v-model="address"
              density="compact" 
              persistent-placeholder
              rows="1" 
              auto-grow
              maxlength="255"
              counter
              variant="outlined"
              :readonly="readonly"
            ></v-textarea>
          </v-col>
      </v-row>
      <v-divider class="my-6" />
      <v-row>
        <v-col cols="12" class="d-flex justify-center">
          <v-btn
            class="mr-4"
            variant="outlined"
            color="primary"
            min-width="120"
            @click="onCancel"
          >
            {{ typeModelIs === 'add' ? 'ยกเลิก' : 'ปิด' }}
          </v-btn>
          <v-btn v-if="typeModelIs === 'add'"
            color="primary"
            min-width="120"
            prepend-icon="mdi-content-save"
            style="color: white"
            @click="onSubmit"
          >
            บันทึก
          </v-btn>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>

</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import apiMember from '../api/member';
import DatePicker from '../compoments/DatePicker.vue';
import Swal from 'sweetalert2'

const name = ref('')
const surname = ref('')
const address = ref('')
const dateOfBirth = ref(null)
const age = ref('')

const dialog = defineModel({ type: Boolean, default: false });
const emits = defineEmits(['success-submit']);

const props = defineProps({
  typeModel: { type: String, default: 'add' },
  readonly: {
    type: Boolean,
    default: false,
  },
  dataMember: {
    type: Object,
    default() {
      return {}
    }
  }

})
const typeModelIs = ref(props.typeModel);
const readonly = ref(props.readonly);

const onCancel = () => {
  dialog.value = false;
}


 const onSubmit = async () => {
  const dataObject = {
    name: name.value,
    surname: surname.value,
    address: address.value,
    dateOfBirth: dateOfBirth.value,
    age: computedAge.value
  };
  // console.log('dataObject :>> ', dataObject);
  try {
    const result = await apiMember.create(dataObject)
    if (result.success) {
      Swal.fire({
        title: result.response.message,
        icon: "success"
      })
      emits('success-submit');
    } else {
      console.error(result.message)
    }
  } catch (error) {
    console.error(error)
  }

}

const computedAge = computed(() => {
  if (!dateOfBirth.value || dateOfBirth.value.length !== 8) return null;

  let year = parseInt(dateOfBirth.value.substring(0, 4), 10);
  const month = parseInt(dateOfBirth.value.substring(4, 6), 10) - 1;
  const day = parseInt(dateOfBirth.value.substring(6, 8), 10);

  // แปลง พ.ศ. เป็น ค.ศ.
  if (year > 2400) year -= 543;

  const today = new Date();

  let age = today.getFullYear() - year;

  const m = today.getMonth() - month;
  if (m < 0 || (m === 0 && today.getDate() < day)) {
    age--;
  }

  return age;
})

// function calculateAge(dateStr) {
//   if (!dateStr || dateStr.length !== 8) return null;

//   let year = parseInt(dateStr.substring(0, 4), 10);
//   const month = parseInt(dateStr.substring(4, 6), 10) - 1;
//   const day = parseInt(dateStr.substring(6, 8), 10);

//   // แปลง พ.ศ. เป็น ค.ศ.
//   if (year > 2400) year -= 543;

//   const today = new Date();

//   let age = today.getFullYear() - year;

//   const m = today.getMonth() - month;
//   if (m < 0 || (m === 0 && today.getDate() < day)) {
//     age--;
//   }

//   return age;
// }


const setDataEdit = async () => {
  name.value = props.dataMember.name ?? ''
  surname.value = props.dataMember.surname ?? ''
  dateOfBirth.value = props.dataMember.dateOfBirth ?? ''
  address.value = props.dataMember.address ?? ''
  age.value = props.dataMember.age ?? ''
}

onMounted(() => {
  if (props.typeModel == "view") {
    setDataEdit()
  }
})

</script>

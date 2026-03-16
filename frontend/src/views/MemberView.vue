<template>
  <div>
    <v-container class="px-4">
      <v-card class="pa-8 pa-sm-10" rounded="xl">
        <v-row justify="end" class="mt-1 mb-2 mr-3">
          <v-col cols="auto">
            <v-btn
              color="primary"
              prepend-icon="mdi-plus"
              style="color: white"
              height="40"
              @click="onOpenDialog()"
            >
              ADD
            </v-btn>
          </v-col>
        </v-row>
          <v-data-table-server
            v-model:items-per-page="perPage"
            :headers="headers"
            :items="items"
            :items-length="totalItems"
            :loading="loading"
            @update:options="fetchMembers"
          >
            <template #item="{ item }">
              <tr>
                <td style="text-align: center;">{{ item.id }}</td>
                <td style="text-align: center;">{{ item.name }} {{ item.surname }}</td>
                <td style="text-align: center;">{{ item.address }}</td>
                <td style="text-align: center;">{{ formatDateToDDMMYYYY(item.dateOfBirth) }}</td>
                <td style="text-align: center;">{{ item.age }}</td>
                <td style="text-align: center;">
                  <v-btn prepend-icon="mdi-text-box-search" size="small" variant="outlined" @click="onSelected(item)">
                    View
                  </v-btn>
                </td>
              </tr>
            </template>
          </v-data-table-server>

      </v-card>


    <v-dialog
      v-model="memberDialog"
      persistent
      width="40vw"
    >
      <menberModals 
        v-model="memberDialog" 
        :type-model="typeModel" 
        :readonly="readonly"
        :data-member="dataMember"
        @success-submit="() => { memberDialog = false; refresh(); }" />
    </v-dialog>

    </v-container>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import apiMember from '../api/member';
import menberModals from '../modals/memberModals.vue';

const memberDialog = ref(false);
const typeModel = ref('');
const readonly = ref(false);
let dataMember = ref([])
const items = ref([]);
const loading = ref(false);
const perPage = ref(10)
const totalItems = ref(0)

const headers = reactive([
  { title: 'id', key: 'id', align: 'center' },
  { title: 'ชื่อ-สกุล', key: 'fullname', align: 'center' },
  { title: 'ที่อยู่', key: 'address', align: 'center' },
  { title: 'วันเกิด', key: 'dateOfBirth', align: 'center' },
  { title: 'อายุ', key: 'age', align: 'center' },
  { title: 'action', key: 'action', align: 'center', sortable: false },
])


const onSelected = (item) => {
  typeModel.value = 'view'
  dataMember.value = item
  readonly.value = true
  memberDialog.value = true;
}

function onOpenDialog() {
  typeModel.value = 'add'
  readonly.value = false
  memberDialog.value = true;
}

function formatDateToDDMMYYYY(date) {
  let data = date.toString().padEnd(8, "0");
  let [year, month, day] = [
    data.substring(0, 4),
    data.substring(4, 6),
    data.substring(6),
  ];
  return `${day}/${month}/${year}`;
}

const fetchMembers = async ({ page = 1, itemsPerPage = 10 } = {}) => {
  const sendData = {
    page: page || 1, // หน้าปัจจุบัน
    limit: itemsPerPage || 10, // จํานวนข้อมูลต่อหน้า
  }
  try {
    const res = await apiMember.getAll(sendData)

    if (!res?.success) {
      throw new Error(res?.message || 'Fetch members failed')
    }

    const payload = res?.response

    items.value = (payload.data || []).map((item, index) => ({
      id: (page - 1) * itemsPerPage + index + 1,
      ...item
    }))

    totalItems.value = payload?.total || 0

  } catch (error) {
    console.error('fetchMembers error:', error)

    items.value = []
    totalItems.value = 0
  }
  //mockData();
}

const clearData = () => {
  perPage.value = 10
  fetchMembers()
}

const refresh = () => {
  clearData();
}

// function mockData() {
//   items.value = {
//     data: [
//       { id: "1", name: "African", surname: "Elephant", address: "Loxodonta africana", dateOfBirth: "25660817", age: "3" },
//       { id: "2", name: "Asian", surname: "Elephant", address: "Elephas maximus", dateOfBirth: "25650512", age: "4" },
//       { id: "3", name: "Bengal", surname: "Tiger", address: "Panthera tigris tigris", dateOfBirth: "25641103", age: "5" },
//       { id: "4", name: "Lion", surname: "Red", address: "Panthera leo", dateOfBirth: "25630225", age: "6" },
//       { id: "5", name: "Cheetah", surname: "veryfast", address: "Acinonyx jubatus", dateOfBirth: "25650914", age: "4" },
//       { id: "6", name: "Leopard", surname: "blackdot", address: "Panthera pardus", dateOfBirth: "25640719", age: "4" },
//       { id: "7", name: "Giraffe", surname: "longneck", address: "Giraffa camelopardalis", dateOfBirth: "25620108", age: "7" },
//       { id: "8", name: "Zebra", surname: "skinblackorwhite", address: "Equus quagga", dateOfBirth: "25660630", age: "2" },
//       { id: "9", name: "Hippo", surname: "potamus", address: "Hippopotamus amphibius", dateOfBirth: "25610311", age: "8" },
//       { id: "10", name: "Rhino", surname: "ceros", address: "Rhinocerotidae", dateOfBirth: "25631022", age: "5" }
//     ],
//     page: 1,
//     limit: 10,
//     total: 20,
//     totalPages: 2
//   }

//   totalItems.value = items.value.total
// }


onMounted(fetchMembers)
</script>

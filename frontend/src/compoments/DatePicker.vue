<template>
  <v-row>
    <v-col cols="12" class="pa-0">
      <v-menu
        v-model="menu"
        :close-on-content-click="false"
        transition="scale-transition"
        location="bottom end"
        :min-width="0"
        :disabled="setDisabled"
      >
        <template #activator="{ props }">
          <v-text-field
            v-bind="props"
            v-model="dateText"
            :label="label"
            append-inner-icon="mdi-calendar"
            :readonly="setDisabled"
          />
        </template>

        <v-sheet width="auto">
          <v-date-picker
            v-model="dateCalendar"
            color="primary"
            hide-actions
            :disabled="setDisabled"
          />
        </v-sheet>
      </v-menu>
    </v-col>
  </v-row>
</template>

<script setup>
import { ref, computed, watch, onMounted } from "vue";

const props = defineProps({
  modelValue: { default: 0 },
  label: { default: "วันที่" },
  disabled: { default: false },
  type: { default: "" }
});

const emit = defineEmits(["update:modelValue"]);

const date = ref(null);
const dateFormat = ref(null);
const menu = ref(false);

const setDisabled = computed(() => props.disabled);

// แปลงจาก JS Date object หรือ "YYYY-MM-DD" → "DDMMYYYY" (พ.ศ.)
const formatDate = (val) => {
  if (!val) return null;
  let dateStr;
  if (val instanceof Date) {
    const y = val.getFullYear();
    const m = String(val.getMonth() + 1).padStart(2, "0");
    const d = String(val.getDate()).padStart(2, "0");
    dateStr = `${y}-${m}-${d}`;
  } else {
    dateStr = val;
  }
  let [year, month, day] = dateStr.split("-");
  return `${day?.padStart(2, "0")}/${month?.padStart(2, "0")}/${String(
    Number(year) + 543
  ).padStart(4, "0")}`;
};

// "DD/MM/YYYY" (พ.ศ.) → ตัวเลข YYYYMMDD (พ.ศ.) สำหรับ emit
const formatToYYYYMMDD = (val) => {
  if (val == null) return 0;
  const str = val.toString();
  // รองรับทั้ง "DD/MM/YYYY" และ "DDMMYYYY"
  const clean = str.replaceAll("/", "");
  let [year, month, day] = [
    clean.substring(4),
    clean.substring(2, 4),
    clean.substring(0, 2)
  ];
  return `${year}${month}${day}`;
};

// ตัวเลข YYYYMMDD (พ.ศ.) → "DD/MM/YYYY" (พ.ศ.)
const formatToDDMMYYYY = (date) => {
  if (date == null) return null;
  let data = `${date.toString().padEnd(8, "0")}`;
  let [year, month, day] = [
    data.substring(0, 4),
    data.substring(4, 6),
    data.substring(6)
  ];
  return `${day}/${month}/${year}`;
};

const getDayMonthYearForZeroPad = (day, month, year) => {
  let y = String(Math.abs(year - 543)).padStart(4, "0");
  let m = Number(month) ? `${month.padStart(2, "0")}` : "00";
  let d = Number(day) ? `${day.padStart(2, "0")}` : "00";
  return [y, m, d];
};

// "DD/MM/YYYY" (พ.ศ.) → "YYYY-MM-DD" (ค.ศ.) สำหรับ v-date-picker
const parseDate = (dateVal) => {
  if (!dateVal || dateVal.length !== 10) return null;
  let day = dateVal.substring(0, 2);
  let month = dateVal.substring(3, 5);
  let year = Number(dateVal.substring(6));
  let [y, m, d] = getDayMonthYearForZeroPad(day, month, year);
  return `${y}-${m}-${d}`;
};

const dateText = computed({
  get() {
    return dateFormat.value;
  },
  set(newVal) {
    if (newVal) {
      dateFormat.value = newVal;
      date.value = parseDate(newVal);
      emit("update:modelValue", formatToYYYYMMDD(newVal));
    } else {
      dateFormat.value = null;
      date.value = null;
      emit("update:modelValue", 0);
    }
  }
});

// Vuetify 3: v-date-picker v-model รับ/คืน Date object
const dateCalendar = computed({
  get() {
    if (!date.value) return null;
    const [y, m, d] = date.value.split("-").map(Number);
    return new Date(y, m - 1, d);
  },
  set(newVal) {
    menu.value = false;
    if (!newVal) return;
    const formatted = formatDate(newVal);
    dateFormat.value = formatted;
    date.value = parseDate(formatted);
    emit("update:modelValue", formatToYYYYMMDD(formatted));
  }
});

const setDate = () => {
  if (!props.modelValue) return;
  date.value = props.modelValue;
  dateText.value = formatToDDMMYYYY(date.value);
};

watch(
  () => props.modelValue,
  (newVal) => {
    if (newVal) {
      if (dateFormat.value == null || dateFormat.value.length === 8) {
        dateText.value = formatToDDMMYYYY(newVal);
      }
    } else {
      dateText.value = null;
    }
  }
);


onMounted(() => {
  setDate();
});
</script>

<style scoped>
.text-no-border :deep(.v-input__control::before),
.text-no-border :deep(.v-input__control::after),
.text-no-border :deep(input) {
  border-style: none;
  margin-bottom: 0px;
}
</style>

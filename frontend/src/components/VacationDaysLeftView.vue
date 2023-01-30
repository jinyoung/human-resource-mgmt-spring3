<template>

    <v-data-table
        :headers="headers"
        :items="vacationDaysLeft"
        :items-per-page="5"
        class="elevation-1"
    ></v-data-table>

</template>

<script>
    const axios = require('axios').default;

    export default {
        name: 'VacationDaysLeftView',
        props: {
            value: Object,
            editMode: Boolean,
            isNew: Boolean
        },
        data: () => ({
            headers: [
                { text: "id", value: "id" },
            ],
            vacationDaysLeft : [],
        }),
          async created() {
            var temp = await axios.get(axios.fixUrl('/vacationDaysLefts'))

            temp.data._embedded.vacationDaysLefts.map(obj => obj.id=obj._links.self.href.split("/")[obj._links.self.href.split("/").length - 1])

            this.vacationDaysLeft = temp.data._embedded.vacationDaysLefts;
        },
        methods: {
        }
    }
</script>


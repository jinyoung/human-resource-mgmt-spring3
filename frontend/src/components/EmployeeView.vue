<template>

    <v-data-table
        :headers="headers"
        :items="employee"
        :items-per-page="5"
        class="elevation-1"
    ></v-data-table>

</template>

<script>
    const axios = require('axios').default;

    export default {
        name: 'EmployeeView',
        props: {
            value: Object,
            editMode: Boolean,
            isNew: Boolean
        },
        data: () => ({
            headers: [
                { text: "id", value: "id" },
            ],
            employee : [],
        }),
          async created() {
            var temp = await axios.get(axios.fixUrl('/employees'))

            temp.data._embedded.employees.map(obj => obj.id=obj._links.self.href.split("/")[obj._links.self.href.split("/").length - 1])

            this.employee = temp.data._embedded.employees;
        },
        methods: {
        }
    }
</script>


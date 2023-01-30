<template>

    <v-data-table
        :headers="headers"
        :items="searchCalendar"
        :items-per-page="5"
        class="elevation-1"
    ></v-data-table>

</template>

<script>
    const axios = require('axios').default;

    export default {
        name: 'SearchCalendarView',
        props: {
            value: Object,
            editMode: Boolean,
            isNew: Boolean
        },
        data: () => ({
            headers: [
                { text: "id", value: "id" },
            ],
            searchCalendar : [],
        }),
          async created() {
            var temp = await axios.get(axios.fixUrl('/searchCalendars'))

            temp.data._embedded.searchCalendars.map(obj => obj.id=obj._links.self.href.split("/")[obj._links.self.href.split("/").length - 1])

            this.searchCalendar = temp.data._embedded.searchCalendars;
        },
        methods: {
        }
    }
</script>


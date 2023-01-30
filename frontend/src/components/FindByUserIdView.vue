<template>

    <v-data-table
        :headers="headers"
        :items="findByUserId"
        :items-per-page="5"
        class="elevation-1"
    ></v-data-table>

</template>

<script>
    const axios = require('axios').default;

    export default {
        name: 'FindByUserIdView',
        props: {
            value: Object,
            editMode: Boolean,
            isNew: Boolean
        },
        data: () => ({
            headers: [
                { text: "id", value: "id" },
            ],
            findByUserId : [],
        }),
          async created() {
            var temp = await axios.get(axios.fixUrl('/findByUserIds'))

            temp.data._embedded.findByUserIds.map(obj => obj.id=obj._links.self.href.split("/")[obj._links.self.href.split("/").length - 1])

            this.findByUserId = temp.data._embedded.findByUserIds;
        },
        methods: {
        }
    }
</script>


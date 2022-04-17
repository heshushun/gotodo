<template>
    <div>
        <h2>cpu 使用量:</h2>
        <p>status: {{cpu.status}}</p> <p>info: {{cpu.info}}</p>

        <h2>ram 使用量:</h2>
        <p>status: {{ram.status}}</p> <p>info: {{ram.info}}</p>

        <h2>disk 使用量:</h2>
        <p>status: {{disk.status}}</p> <p>info: {{disk.info}}</p>
    </div>
</template>

<script>
    import axios from 'axios'

    export default {
        name: "Health",
        data() {
            return {
                cpu: "",
                ram: "",
                disk: ""
            }
        },
        methods: {
            getHealth() {
                axios.all([
                    axios.get("/api/sd/cpu"),
                    axios.get("/api/sd/disk"),
                    axios.get("/api/sd/ram"),

                ]).then(axios.spread((cpu, disk, ram) => {
                    this.cpu = cpu.data;
                    this.disk = disk.data;
                    this.ram = ram.data;
                })).catch((error) => {
                    console.error(error);
                });
            }
        },

        mounted() {
            this.getHealth()
        }

    }
</script>

<style scoped>

</style>
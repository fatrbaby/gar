const { createApp, ref, reactive, onMounted } = Vue

createApp({
    setup() {
        const form = reactive({
            keyword: "",
            author: "",
            viewCount: 0,
            categories: [],
        })

        const categories = ref([])

        onMounted(async () => {
            try {
                const r = await fetch('/api/categories')
                categories.value = await r.json()
            } catch (e) {
                console.log(e)
            }
        });

        return {
            form,
            categories,
        }
    }
}).mount('#app')
const {createApp, ref, reactive, onMounted} = Vue

createApp({
    setup() {
        const form = reactive({
            keyword: "",
            author: "",
            viewCount: 0,
            viewsFrom: 10,
            viewsTo: 100000,
            categories: [],
        })

        const videos = ref([])
        const categories = ref([])

        onMounted(async () => {
            try {
                const r = await fetch('/api/categories')
                categories.value = await r.json()
            } catch (e) {
                console.log(e)
            }
        });

        const doSearch = async () => {
            try {
                const data = {
                    author: form.author,
                    categories: form.categories,
                    keywords: form.keyword.split(',').filter(Boolean), // 过滤空关键字
                    viewsFrom: form.viewsFrom,
                    viewsTo: form.viewsTo,
                }

                const r = await fetch("/api/search", {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(data),
                })

                if (r.ok) {
                    videos.value = await r.json()
                    console.log(videos.value)
                } else {
                    const errorMessage = await r.text(); // 获取错误消息
                    console.log(`HTTP error! status: ${r.status}, body: ${errorMessage}`)
                }
            } catch (e) {
                console.log(e.message); // 捕获并打印异常
            }
        }

        const dateFormat = (ts) => {
            return new Date(ts * 1000).toISOString().split('T')[0];
        }

        return {
            form,
            categories,
            videos,
            doSearch,
            dateFormat,
        }
    }
}).mount('#app')
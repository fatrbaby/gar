const {createApp, ref, reactive, onMounted} = Vue

createApp({
    setup() {
        const form = reactive({
            keyword: "",
            author: "",
            viewsFrom: 0,
            viewsTo: 0,
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
                    keywords: form.keyword.split(',').filter(Boolean),
                    viewsFrom: Number(form.viewsFrom),
                    viewsTo: Number(form.viewsTo),
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
            if (typeof ts !== 'number' || isNaN(ts)) {
                return 'error time'
            }

            const t = new Date(ts * 1000)
            const y = t.getFullYear()
            const m = String(t.getMonth() + 1).padStart(2, '0')
            const d = String(t.getDate()).padStart(2, '0')
            const h = String(t.getHours()).padStart(2, '0')
            const i = String(t.getMinutes()).padStart(2, '0')
            const s = String(t.getSeconds()).padStart(2, '0')

            return `${y}-${m}-${d} ${h}:${i}:${s}`;
        };

        return {
            form,
            categories,
            videos,
            doSearch,
            dateFormat,
        }
    }
}).mount('#app')
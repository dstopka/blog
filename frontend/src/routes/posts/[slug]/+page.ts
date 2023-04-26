/** @type {import('./$types').PageLoad} */

export async function load({ fetch, params }: { fetch: any, params: any}) {
    const resp = await fetch(`http://localhost:8001/posts?slug=${params.slug}`);
    const postData = await resp.json()

    return {
        post: postData
    }
}
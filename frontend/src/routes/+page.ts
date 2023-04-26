/** @type {import('./$types').PageLoad} */

export async function load({ fetch }: { fetch: any }) {
    const res = await fetch('http://localhost:8008/posts')
    const postsResp = await res.json()

    return postsResp
}
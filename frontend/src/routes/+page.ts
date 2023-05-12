import { PUBLIC_ALL_POSTS_ADDR } from '$env/static/public';

/** @type {import('./$types').PageLoad} */

export async function load({ fetch }: { fetch: any }) {
    const res = await fetch(`${PUBLIC_ALL_POSTS_ADDR}/posts`)
    const postsResp = await res.json()

    return postsResp
}
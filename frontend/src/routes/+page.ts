import { PUBLIC_ALL_POSTS_ADDR } from '$env/static/public';

/** @type {import('./$types').PageLoad} */
export async function load({ fetch, setHeaders }: { fetch: any, setHeaders: any }) {
    const res = await fetch(`${PUBLIC_ALL_POSTS_ADDR}/posts`)
    const postsResp = await res.json()

    setHeaders({
        'cache-control': 'public, max-age=300'
    });

    return postsResp
}
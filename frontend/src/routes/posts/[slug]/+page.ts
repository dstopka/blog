import { PUBLIC_POST_HOST } from '$env/static/public';

/** @type {import('./$types').PageLoad} */

export async function load({ fetch, params }: { fetch: any, params: any}) {
    const resp = await fetch(`${PUBLIC_POST_HOST}/post?slug=${params.slug}`);
    const postData = await resp.json()

    return {
        post: postData
    }
}
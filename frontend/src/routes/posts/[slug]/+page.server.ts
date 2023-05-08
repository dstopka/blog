import { PUBLIC_POST_HOST } from '$env/static/public';
import { compile } from 'mdsvex';
import remarkBreaks from 'remark-breaks';

/** @type {import('./$types').PageLoad} */

export async function load({ fetch, params }: { fetch: any, params: any }) {
    const resp = await fetch(`${PUBLIC_POST_HOST}/post?slug=${params.slug}`);
    const postData = await resp.json()

    console.log(postData.content)

    const compiledContent = (await compile(postData.content, {
        remarkPlugins: [remarkBreaks]
    }))?.code
        .replace(/>{@html `<code class="language-/g, '><code class="language-')
        .replace(/<\/code>`}<\/pre>/g, '</code></pre>');
    postData.content = compiledContent;

    return {
        post: postData
    }
}
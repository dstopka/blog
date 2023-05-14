<script lang="ts">
	import Post from '$lib/components/Post.svelte';

	export let data;
	const { post } = data;
</script>

<svelte:head>
	<meta property="og:title" content={post.title} />
	<title>{post.title}</title>
</svelte:head>

<Post post={post}/>

<style>
</style>

<script lang="ts">
	import HeaderLine from '$lib/HeaderLine/HeaderLine.svelte';
	import PostsList from '$lib/Post/List.svelte';

	/** @type {import('./$types').PageData} */
	export let data: any;

	const { posts }: { posts: any[] } = data;
</script>

<header>
	<HeaderLine header="All Posts" size="3rem" />
</header>
<PostsList {posts} />

<style>
	header {
		max-width: 55rem;
		margin: auto;
		padding: 0 1.5rem;
	}
</style>

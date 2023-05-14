<script lang="ts">
	import HeaderLine from '$lib/components/HeaderLine.svelte';
	import PostsList from '$lib/components/PostsList.svelte';

	/** @type {import('./$types').PageData} */
	export let data: any;

	const { posts }: { posts: any[] } = data;
</script>

<header>
	<HeaderLine header="All Posts" />
</header>
<PostsList {posts} />

<style>
	header {
		font-size: 2.5rem;
		max-width: 55rem;
		margin: 1rem auto;
		padding: 0 1.5rem;
	}

	@media (max-width: 768px) {
		header {
			font-size: 2rem;
		}
	}
</style>

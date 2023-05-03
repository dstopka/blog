<script lang="ts">
	import Pagination from '$lib/Pagination/Pagination.svelte';
	import Summary from '$lib/Post/Summary.svelte';

	/** @type {import('./$types').PageData} */
	export let data: any;

	const { posts }: { posts: any[] } = data;

	let postsToDisplay: any[] = [];
</script>

<section class="posts-list">
	{#each postsToDisplay as post}
		<Summary {post} />
	{/each}
	<Pagination allPosts={posts} bind:postsToDisplay />
</section>

<style>
	.posts-list {
		flex-direction: column;
		display: flex;
	}
</style>

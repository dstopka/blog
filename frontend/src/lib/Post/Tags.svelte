<script lang="ts">
	export let tags: any[];
</script>

<div class="tags">
	{#each tags as tag}
		<span class="tag">{tag}</span>
	{/each}
</div>

<style>
	.tags {
		display: flex;
		flex-wrap: wrap;
		row-gap: 0.5rem;
        column-gap: 0.5rem;
		align-items: center;
	}

	.tag {
		padding: 0.2rem 0.6rem;
        background-color: var(--color-bg-gray);
		border-radius: 0.4rem;
		color: var(--color-text-dark-gray);
		/* color: #FFF; */
		font-weight: 400;
        font-size: 0.9rem;
	}
</style>

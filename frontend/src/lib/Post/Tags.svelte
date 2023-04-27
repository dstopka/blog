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
        background-color: rgb(236, 236, 236);
		border-radius: 0.6rem;
		color: rgb(74, 74, 74);
		font-weight: 400;
        font-size: 0.9rem;
	}
</style>

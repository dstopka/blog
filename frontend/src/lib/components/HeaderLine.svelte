<script lang="ts">
	export let header: string;
	export let size: string;
</script>

<div class="header-line">
	<h2 style="font-size: {size}">{header}</h2>
	<hr />
</div>

<style>
	.header-line {
		margin: 2rem 0 1rem 0;
		display: flex;
		align-items: center;
		gap: 1rem;
        color: var(--color-text-dark-headings);
	}

	hr {
		flex-grow: 1;
		background-color: var(--color-logo-secondary-bg);
		height: 1px;
		border: none;
	}
</style>

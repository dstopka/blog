<script lang="ts">
	export let header: string;
</script>

<div class="header-line">
	<h2>{header}</h2>
	<hr />
</div>

<style>
	.header-line {
		display: flex;
		align-items: center;
		gap: 1rem;
        color: var(--color-text-dark-headings);
	}

	h2 {
		font-size: inherit;
	}

	hr {
		flex-grow: 1;
		background-color: var(--color-logo-secondary-bg);
		height: 1px;
		border: none;
	}
</style>

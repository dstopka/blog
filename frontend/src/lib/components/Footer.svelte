<script>
	import Socials from '$lib/components/Socials.svelte';
</script>

<footer>
	<div class="container">
		<div class="footer-legal-content">
			<span> &copy; Copyright 2023 Dariusz Stopka </span>
		</div>
		<div class="socials-list">
			<Socials
				color={'var(--color-text-light-primary)'}
				colorHover={'var(--color-text-light-primary)'}
				inCircle={true}
				gap={'1rem'}
			/>
		</div>
	</div>
</footer>

<style>
	footer {
		margin-top: auto;
		background-color: var(--color-footer-bg);
	}
	
	.container {
		display: flex;
		flex-direction: row;
		align-items: center;
		max-width: 70rem;
		padding: 0 1rem;
		margin: auto;
	}

	.socials-list {
		margin-left: auto;
		padding: 1rem 0;
		height: 100%;
		font-size: 1.25rem;
	}
	
	.footer-legal-content {
		font-size: 0.75rem;
		color: var(--color-text-light-secondary);
		margin-right: auto;
	}

	@media (max-width: 768px) {
		.footer-legal-content {
			margin: auto;
		}

		.socials-list {
			margin: auto;
		}
	}
</style>

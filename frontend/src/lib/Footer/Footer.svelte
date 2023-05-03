<script>
	import Socials from '$lib/Socials/Socials.svelte';
</script>

<footer class="page-footer">
	<section class="footer-main">
		<div class="container row">
			<div class="socials-list">
				<Socials color={'var(--color-text-light-primary)'} colorHover={'var(--color-text-light-primary)'} inCircle={true} gap={'1rem'} />
			</div>
		</div>
	</section>
	<section class="footer-legal">
		<div class="container row">
			<div class="footer-legal-content">
				<span> &copy; Copyright 2023 Dariusz Stopka </span>
			</div>
		</div>
	</section>
</footer>

<style>
	.page-footer {
		margin-top: auto;
		background-color: var(--color-footer-bg);
	}

	.container {
		max-width: 60rem;
		padding: 0 1rem;
		margin: auto;
	}

	.row {
		display: flex;
		flex-direction: row;
		align-items: center;
		height: 100%;
	}

	.footer-main {
		font-size: 1.5rem;
		padding: 0 1rem;
		width: 100%;
	}

	.socials-list {
		margin-left: auto;
		padding: 1rem 0;
		height: 100%;
	}

	.footer-legal {
		padding-bottom: 1rem;
		font-size: 0.8rem;
		color: var(--color-text-light-secondary);
	}

	.footer-legal-content {
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

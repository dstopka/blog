<script>
	import Socials from '$lib/Socials/Socials.svelte';
</script>

<footer class="page-footer">
	<section class="footer-main">
		<div class="footer-main-container">
			<div class="socials-list">
				<Socials color={'light'} />
			</div>
		</div>
	</section>
	<section class="footer-legal">
		<div class="footer-legal-content">
			<span> &copy; Copyright 2023 Dariusz Stopka </span>
		</div>
	</section>
</footer>

<style>
	.page-footer {
		color: rgb(252, 252, 252);
		background: linear-gradient(180deg, rgb(56, 56, 56) 0%, rgb(34, 34, 34) 100%);
	}

	.footer-main {
		height: 50px;
		line-height: 50px;
	}

	.socials-list {
		margin: auto;
	}

	.footer-main-container {
		max-width: 60rem;
		font-size: 1.5rem;
		height: 100%;
		display: flex;
		flex-direction: row;
		margin: auto;
	}

	.footer-legal {
		padding: 0 1rem;
		height: 50px;
		line-height: 50px;
		font-size: 0.8rem;
	}

	.footer-legal-content {
		text-align: left;
	}

	@media (max-width: 768px) {
		.footer-legal-content {
			text-align: center;
		}
	}
</style>

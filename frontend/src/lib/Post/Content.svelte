<script lang="ts">
	export let content: any;
</script>

<svelte:head>
	<link
		rel="stylesheet"
		href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/styles/monokai.min.css"
	/>
</svelte:head>

<section class="post-content">
	{@html content}
</section>

<style>
	section {
		margin: 0.5rem 0;
	}

	/* Headings */

	:global(.post-content h1, .post-content h2, .post-content h3) {
		color: var(--color-text-dark-headings);
		margin: 1rem 0;
	}

	:global(.post-content h1) {
		font-size: 1.75rem;
	}

	/* Tasks list */

	:global(.post-content li:has(input)) {
		list-style-type: none;
	}

	/* Lists */

	:global(ul, ol) {
		padding: 0px 25px;
	}

	/* Image */

	:global(.post-content img.content-image) {
		display: block;
		margin-left: auto;
		margin-right: auto;
	}

	/* Code */

	:global(.post-content code) {
		border-radius: 0.2rem;
	}

	:global(.post-content code) {
		font-size: 1rem;
		line-height: 1.5rem;
		font-family: 'Ubuntu Mono', monospace;
		tab-size: 4;
	}

	:global(.post-content[class*='hljs-']) {
		font-weight: 400 !important;
	}

	:global(pre) {
		margin: 1.5rem 0;
	}

	/* Quote */

	:global(.post-content blockquote) {
		background-color: var(--color-bg-gray);
		padding: 0.5rem 1rem;
		margin: 2rem 0;
		line-height: normal;
		border-radius: 0.2rem;
		font-weight: 300;
		font-size: 1rem;
	}

	:global(.post-content blockquote p) {
		margin: 0.5rem 0;
		color: var(--color-text-dark-gray);
	}

	/* Captions */

	:global(.post-content figcaption) {
		text-align: center;
		margin-bottom: 1.5rem;
	}

	:global(.post-content figcaption p) {
		margin-top: 0;
		font-size: 0.8rem;
	}

	:global(.post-content pre:has(+ figcaption), .post-content img:has(+ figcaption)) {
		margin-bottom: 0.5rem;
	}

	/* Links */

	:global(.post-content a) {
		color: var(--color-text-dark-decorator);
		text-decoration: none;
		font-weight: bold;
	}

	/* Divider */

	:global(hr) {
		margin: 2rem 0;
		color: var(--color-bg-gray);
	}
</style>

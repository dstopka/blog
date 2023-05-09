<script lang="ts">
	export let content: any;
</script>

<section class="post-content">
	{@html content}
</section>

<style>
	:global(.post-content :where(p, pre, hr)) {
		margin-top: 1.5rem;
		margin-bottom: 1.5rem;
	}

	:global(.post-content :where(h1+*, h2+*, h3+*)) {
		margin-top: 0;
	}

	:global(.post-content p:has(+ul)) {
		margin-bottom: 0;
	}

	:global(.post-content > p:last-of-type) {
		margin-bottom: 0;
	}

	/* Headings */

	:global(.post-content h1, .post-content h2, .post-content h3) {
		color: var(--color-text-dark-headings);
		text-align: left;
	}

	:global(.post-content h1) {
		font-size: 1.75rem;
	}

	:global(.post-content h2) {
		font-size: 1.5rem;
	}

	:global(.post-content h3) {
		font-size: 1.25rem;
	}

	:global(.post-content :where(h1, h2, h3)) {
		margin-top: 2rem;
		margin-bottom: 1rem;
	}

	/* Tasks list */

	:global(.post-content li:has(input)) {
		list-style-type: none;
	}

	/* Lists */

	:global(.post-content :where(ul, ol)) {
		padding: 0px 25px;
	}

	/* Image */

	:global(.post-content img) {
		display: block;
		margin-left: auto;
		margin-right: auto;
	}

	/* Code */
	:global(.post-content p code) {
		font-family: 'Fira Mono', monospace;
		color: var(--color-text-dark-decorator);
		background-color: var(--color-bg-gray);
		font-size: 0.875rem;
		padding: 0.125rem 0.25rem;
		border-radius: 0.25rem;
	}

	:global(.post-content pre:has(code)) {
		border-radius: 0.5rem;
	}

	:global(.post-content pre code[class*='language-']) {
		font-family: 'Fira Mono', monospace;
		font-size: 0.875rem;
		line-height: 1.5rem;
		display: inline-block;
	}

	/* Quote */

	:global(.post-content blockquote) {
		background-color: var(--color-bg-gray);
		padding: 0.5rem 1rem;
		line-height: normal;
		border-radius: 0.2rem;
		font-weight: 300;
		font-size: 1rem;
	}

	:global(.post-content blockquote p) {
		color: var(--color-text-dark-gray);
		margin: 0.5rem 0;
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

	:global(.post-content pre:has(+ figcaption), .post-content p:has(+ figcaption)) {
		margin-bottom: 0.25rem;
	}

	/* Links */

	:global(.post-content a) {
		color: var(--color-text-dark-decorator);
		text-decoration: none;
		font-weight: bold;
	}

	/* Divider */

	:global(.post-content hr) {
		background-color: var(--color-logo-secondary-bg);
		height: 1px;
		border: none;
	}

</style>

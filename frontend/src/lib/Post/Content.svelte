<script lang="ts">
	import { marked } from 'marked';
	import hljs from 'highlight.js';

	export let content: string;

	marked.setOptions({
		highlight: function (code, lang) {
			const language = hljs.getLanguage(lang) ? lang : 'plaintext';
			return hljs.highlight(code, { language }).value;
		},
		langPrefix: 'hljs language-',
		gfm: true,
		breaks: true
	});

	marked.Renderer.prototype.paragraph = (text) => {
		if (text.startsWith('<img')) {
			return text + '\n';
		}
		return '<p>' + text + '</p>';
	};

	marked.Renderer.prototype.image = (text) => {
		return `<img class="content-image" src="` + text + `"/>`;
	};
</script>

<svelte:head>
	<link
		rel="stylesheet"
		href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/styles/monokai.min.css"
	/>
</svelte:head>

<section class="post-content">
	<div>{@html marked(content)}</div>
</section>

<style>
	.post-content {
		margin: 0.5rem 0;
		line-height: 1.7;
	}

	/* Text */

	:global(.post-content p) {
		font-size: 1.1rem;
	}

	/* Headings */

	:global(.post-content h1, .post-content h2, .post-content h3) {
		color: var(--color-text-dark-headings);
	}

	:global(.post-content h1) {
		font-size: 1.8rem;
		font-weight: 700;
		margin-block-start: 0.5em;
		margin-block-end: 0.5em;
	}

	/* Taks list */

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
	}

	:global(.post-content blockquote:before) {
		font-family: "Font Awesome 5 Free";
		content: "\f10d";
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
	}

	/* Divider */

	:global(hr) {
		margin: 2rem 0;
		color: var(--color-bg-gray);
	}
</style>

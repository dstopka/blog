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
		return `<img class="content-image" src="` + text + `""/>`;
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
	}

	/* Headings */

	:global(.post-content h1) {
		font-size: 1.6rem;
		font-weight: 700;
		margin-block-start: 0.5em;
		margin-block-end: 0.5em;
	}

	/* Taks list */

	:global(.post-content li:has(input)) {
		list-style-type: none;
	}

	/* Image */

	:global(.post-content img.content-image) {
		display: block;
		margin-left: auto;
		margin-right: auto;
	}

	/* Code */

	:global(.post-content code) {
		border-radius: 0.5rem;
	}

	:global(.post-content code) {
		font-family: Consolas, monaco, monospace;
		font-size: 0.9rem;
	}

	:global(.post-content[class*='hljs-']) {
		font-weight: 400 !important;
	}

	/* Quote */

	:global(.post-content blockquote) {
		background-color: #f4f3f1;
		border-left: 5px solid #dbdbdb;
		padding: 0.5rem 1rem;
		margin: 2rem 0;
        font-size: 1rem;
	}

    :global(.post-content blockquote p) {
		margin: 0.5rem 0;
        color: rgb(74, 74, 74);
	}

	/* Captions */

	:global(.post-content figcaption) {
		text-align: center;
		font-size: 0.8rem;
	}

	:global(.post-content figcaption p) {
		margin-top: 0;
	}

	:global(.post-content pre:has(+ figcaption)) {
		margin-bottom: 0.25rem;
	}

	/* Links */

	:global(.post-content a) {
		color: rgb(95, 187, 224);
	}
</style>

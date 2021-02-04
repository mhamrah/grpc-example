import App from './App.svelte';
import "../public/global.css";

const app = new App({
    target: document.body,
    props: {
        title: 'gRPC Example'
    }
});

export default app;

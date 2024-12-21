

//React file vs open with server 


In React, there are certain features that might not work as expected when opening a file directly in a browser instead of running it through a server.One common reason for this is related to the same-origin policy enforced by web browsers.

The same-origin policy is a security measure that restricts web pages from making requests to a different domain than the one that served the web page. When you open a file directly from the file system using the "file://" protocol, the origin is considered to be "null" or "file://", which can lead to issues with certain features in React.

==> Throttling and debouncing 
==> We can not use useState inside class component.
----In React, the useState hook is designed to be used in functional components, not in class components. 
    The useState hook is part of the React Hooks API, which was introduced in React 16.8 to allow functional 
    components to have state and other features previously available only in class components.
    
    
==> We can send html inside props using this
    function component : props.children
    class componenent  : this.props.children
    
==>

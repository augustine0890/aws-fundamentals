# Objectives
- Use Amazon Cloudfront to create Content Delivery Network
- Test your application with Amazon API Gateway

## Amazon CloudFront
- Web services that speeds up distribution of your static and dynamic web content, such as `.html`, `.css`, `.js`, and image files.
- It deliver your content through a worldwide network of data centers (`edge locations`).
- When a user requests content that you're serving with CloudFront --> the user is routed to the edge location that provides the lowest latency.
  - If the content is already in the edge location with the lowest latencty, CloudFront delivers it immediately.
  - If the content is not in that edge location, CloudFront retrieves it from an orgin that you're defined.

## API Gateway
- APIs act as the "front door" for application to access data, business logic, or functionality from your backend services.
- API Gateway handles all the tasks involved in accepting and processing up to hundreds of thousands of concurrent API calls, including traffic management, CORS support, authorization and access control, monitoring, and API version management.
# Massive Pre-Interview Screening

Hey, guys! I'm Vlad and I tried to answer the questions from you.
This repo will contain a code for the Q1 and below I will 
also answer Q2 & Q3.

### Q1 - Coding

See code. The idea is simple - we have some endpoint that can get
cluster loading data for us. So, I used concurrent  loading to get all data faster.

Service interface has like base service interface with one method `Name()` and
Clusters Service should implement base one and has its own methods.

### Q2 - Suppose you have the high-level task of refactoring a module within a project to preserve the existing functionality and add new features. What steps will you take to plan and implement this change?

As for as I understood, I need to split my work from regular development process
and create a new branch for the task. That gives me an opportunity to
no worries about my changes and its impact for the project at all.

#### Steps

1. Make sure that all requirements are understandable
2. Estimate task
3. Create new git branch for the task
4. Encapsulation of public things (vars, funcs, methods) that should be private
5. DRY principle - Don't Repeat Yourself
6. If we have a requirements to update some logic, so we need to make this update.
7. If step 6 exists - need to be sure that all kind of written tests passes
8. If step 6 exists - add new testcases for new logic
9. Manual testing if possible
10. Create new merge request to `main`
11. Make sure that init pipeline is green (build, unit tests)
12. Ask for code review

### Q3 - You have a set of clients that send telemetry messages in JSON format. The general load is 100-200 QPS, with possible spikes up to 5x. Build a solution using AWS resources to receive, pre-process, and store events. You don't need to write code for handling, focus on designing and explaining the usage of resources.

I know that AWS has a service named `Amazon Kinesis Data Streams` that allows
us to handle large amount of streaming data like telemetry messages.
That's about receiving messages.

On the next step we may want to pre-process data somehow. For these needs
I propose to use `AWS Lambda's`. That is serverless compute service and
it could be triggered by Kinesis stream. We can do a chain of some processes
like validation, data transformation, filtering etc.

After pre-processing steps we can store the data using `Amazon Kinesis Data Firehose`.
Using it we can store our data to some scalable and durable storage like S3.
Using S3 we can organize the data by date or any other relevant partitioning strategy.

Of course, to be on wave with services states we need to have some logs.
`Cloud Watch` can help us with that. Also, I can set up alarms to be notified
about any anomalies or issues.

If we need some cache globally (CDN) AWS have `Cloud Front` service for that.



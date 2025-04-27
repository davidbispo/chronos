DONE:

1. Set up the db connection
1. Set up gorm. Using automigrations
1. Set up project root
1. Set up models
1. Set up routes and logic to create appointments, atendees and link them.
1. Set up test suite. Discovered default suite sucks. Moving to testify
1. Discovered automigrations suck on production. 

TODO 

1. Add a proper migration tool. 

1. Add created_at and updated at

1. Add deleted at logic 

1. Find an appointment grouping logic(i.e.: appt series). Look up limitations on that
* Compatibility with fixed schedule and flexible schedule
* Should the "series" have an ical specification. are there any other types of recurrency descriptors
like RRule?
* Should the appointment have an RRULE specification that changes? Should that go to ES? Or stay on MySQL
* Series Rrule logic when rescheduling
* Series generator from Rrule. 

2. Create series database and logic
* Create series: creates child appointments
* Update series: updates Rrule. recreates future appointments. 
* Series update and conflict detection
* Cancel series: cancels future appointments & soft delete
* Cancel reasons

1. Update appointment. Check for conflicts. Rescheduling reason
2. Update attendee.
3. RSVP for event.
* Create a confirmation status table and make the confirmation status an enum. Different types of session
* Explore differnt types of implementation using go for this
* Explore rule for 24h cancellation

4. Create remove attendee. If attendee removed, sessions links removed
5. Create cancel appointment. If appointment removed, session links removed and soft deletion.
* Soft delete for cancelling sessions.

6. Create views for appointments to be able to fetch data
* Appointments by attendant
* All appointments from multiple attendants during a period. Filterable by metadata
* Attendants per appointment with confirmation status

7. Add environment files for test and dev for now. Make test dbs be different
8. Create seeds file and base data for development: check if gorm has seeds

9. Add gore instructions to README.

10. Connect to elasticsearch
* Explore differnt types of implementation using go for db persistence in general
11. Create elasticsearch schema files and add to initdb routine
* Use templates for sharding on appointments per start date?
* Search will always be per start date or id :think

12. Save appointments to elasticsearch.

14. Bulk inserts to db and ES. Scroll api 
15. Conflict detection
16. Rules for ranking better available times
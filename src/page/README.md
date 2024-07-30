# Request Handling (Single Page Application, spa)

This package handles request to the EduHITZone website. 
It is based on HATEOAS principles and uses HTMX to power it.
## Notes:
- Variables use snake_case.
- The `hitdb.Course` has different field structure than `courseItem`, use `toCourseItem` for conversion.
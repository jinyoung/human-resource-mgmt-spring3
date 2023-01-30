package human.resource.mgmt.domain;

import human.resource.mgmt.domain.*;
import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

@RepositoryRestResource(collectionResourceRel = "calendars", path = "calendars")
public interface CalendarRepository
    extends PagingAndSortingRepository<Calendar, String> {}

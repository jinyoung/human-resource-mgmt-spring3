package human.resource.mgmt.domain;

import human.resource.mgmt.domain.*;
import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

@RepositoryRestResource(collectionResourceRel = "vacations", path = "vacations")
public interface VacationRepository
    extends PagingAndSortingRepository<Vacation, String> {}

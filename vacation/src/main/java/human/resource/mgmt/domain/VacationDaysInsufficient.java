package human.resource.mgmt.domain;

import human.resource.mgmt.domain.*;
import human.resource.mgmt.infra.AbstractEvent;
import java.util.*;
import lombok.*;

@Data
@ToString
public class VacationDaysInsufficient extends AbstractEvent {

    private Long id;
    private String vacationId;

    public VacationDaysInsufficient(VacationDaysLeft aggregate) {
        super(aggregate);
    }

    public VacationDaysInsufficient() {
        super();
    }
}

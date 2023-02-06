package human.resource.mgmt.domain;

import human.resource.mgmt.domain.*;
import human.resource.mgmt.infra.AbstractEvent;
import java.util.*;
import lombok.*;

@Data
@ToString
public class VacationDaysIntialized extends AbstractEvent {

    private String userId;
    private Integer dayCount;

    public VacationDaysIntialized(VacationDaysLeft aggregate) {
        super(aggregate);
    }

    public VacationDaysIntialized() {
        super();
    }
}

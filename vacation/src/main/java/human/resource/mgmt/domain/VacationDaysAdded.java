package human.resource.mgmt.domain;

import human.resource.mgmt.domain.*;
import human.resource.mgmt.infra.AbstractEvent;
import java.util.*;
import lombok.*;

@Data
@ToString
public class VacationDaysAdded extends AbstractEvent {

    private String userId;
    private Integer dayCount;
    private String reason;

    public VacationDaysAdded(VacationDaysLeft aggregate) {
        super(aggregate);
    }

    public VacationDaysAdded() {
        super();
    }
}

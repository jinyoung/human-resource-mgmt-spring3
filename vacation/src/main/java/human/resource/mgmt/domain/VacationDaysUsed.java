package human.resource.mgmt.domain;

import human.resource.mgmt.domain.*;
import human.resource.mgmt.infra.AbstractEvent;
import java.util.*;
import lombok.*;

@Data
@ToString
public class VacationDaysUsed extends AbstractEvent {

    private String userId;
    private Integer dayCount;
    private String reason;

    public VacationDaysUsed(VacationDaysLeft aggregate) {
        super(aggregate);
    }

    public VacationDaysUsed() {
        super();
    }
}

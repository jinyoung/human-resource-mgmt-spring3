package human.resource.mgmt.domain;

import human.resource.mgmt.domain.*;
import human.resource.mgmt.infra.AbstractEvent;
import java.util.*;
import lombok.*;

@Data
@ToString
public class VacationApproved extends AbstractEvent {

    private String id;
    private Date startDate;
    private Date endDate;
    private String reason;

    public VacationApproved(Vacation aggregate) {
        super(aggregate);
    }

    public VacationApproved() {
        super();
    }
}
